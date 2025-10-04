// Package fpbxdel deletes FreePBX extensions by invoking a tiny PHP helper
// that calls BMO (Core->delDevice / delUser). No GraphQL required.
package delexts

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	
	"strings"
	"sync"
	"time"
)

const phpHelper = `#!/usr/bin/env php
<?php
if (!isset($argv[1])) { fwrite(STDERR, "Usage: extdel.php <ext>\n"); exit(2); }
$ext = $argv[1];
include '%FREEPBX_CONF%';
$FreePBX = FreePBX::Create();
$device = $FreePBX->Core->getDevice($ext);
$user   = $FreePBX->Core->getUser($ext);
if (is_array($device) && !empty($device["user"])) {
  $FreePBX->Core->delDevice($ext);
  $FreePBX->Core->delUser($device["user"]);
  echo "OK: device+user deleted\n";
} elseif ($user) {
  $FreePBX->Core->delUser($ext);
  echo "OK: user deleted\n";
} else {
  echo "NF: not found\n";
}
`


// Options configures deletion behavior.
type Options struct {
	// Path to php binary (default "php")
	PHPPath string
	// Path to /etc/freepbx.conf (default "/etc/freepbx.conf")
	ConfPath string

	// If true, runs `fwconsole reload` once at the end if all deletions succeeded.
	Reload bool

	// Max parallel deletions (default 1). Keep small; BMO is not heavily concurrent.
	Parallel int

	// Optional logger; if nil, logging is suppressed.
	Logger *log.Logger

	// Timeout for each PHP helper invocation (default 15s if zero).
	PerCallTimeout time.Duration

	// Path to a writable dir for temp files; if empty uses OS default.
	TempDir string

	// ----- fwconsole controls -----

	// Absolute path to fwconsole (default "fwconsole"; commonly "/usr/sbin/fwconsole").
	FwconsolePath string

	// Optional args inserted BEFORE "reload" (e.g. sudo wrapper):
	// Example: []string{"-u","asterisk","/usr/sbin/fwconsole"}
	FwconsoleArgs []string

	// Timeout specifically for the fwconsole reload (defaults to PerCallTimeout if zero).
	ReloadTimeout time.Duration
}


// Result captures the outcome per extension.
type Result struct {
	Ext    string // input extension
	Output string // stdout/stderr combined (trimmed)
	Err    error  // nil on success
}

// Delete deletes all given extensions. It creates a temporary PHP helper,
// invokes it once per extension (optionally in parallel), and optionally runs
// `fwconsole reload` at the end if all succeeded and opts.Reload is true.
func Delete(ctx context.Context, exts []string, opts Options) ([]Result, error) {
	if len(exts) == 0 {
		return nil, errors.New("no extensions provided")
	}
	if opts.PHPPath == "" {
		opts.PHPPath = "php"
	}
	if opts.ConfPath == "" {
		opts.ConfPath = "/etc/freepbx.conf"
	}
	if opts.Parallel <= 0 {
		opts.Parallel = 1
	}
	if opts.PerCallTimeout <= 0 {
		opts.PerCallTimeout = 15 * time.Second
	}

	helperPath, cleanup, err := writeHelper(opts.TempDir, opts.ConfPath)
	if err != nil {
		return nil, fmt.Errorf("write helper: %w", err)
	}
	defer cleanup()

	type job struct{ ext string }
	in := make(chan job)
	var wg sync.WaitGroup

	results := make([]Result, len(exts))
	var idxMu sync.Mutex
	nextIdx := 0

	worker := func() {
		defer wg.Done()
		for j := range in {
			childCtx, cancel := context.WithTimeout(ctx, opts.PerCallTimeout)
			out, e := runPHP(childCtx, opts.PHPPath, helperPath, j.ext)
			cancel()

			idxMu.Lock()
			i := nextIdx
			nextIdx++
			idxMu.Unlock()

			if i < len(results) {
				results[i] = Result{Ext: j.ext, Output: strings.TrimSpace(out), Err: e}
			}
			if opts.Logger != nil {
				if e != nil {
					opts.Logger.Printf("ext %s: ERROR: %v\nOutput: %s\n", j.ext, e, out)
				} else {
					opts.Logger.Printf("ext %s: %s\n", j.ext, out)
				}
			}
		}
	}

	wg.Add(opts.Parallel)
	for w := 0; w < opts.Parallel; w++ {
		go worker()
	}
	for _, e := range exts {
		in <- job{ext: e}
	}
	close(in)
	wg.Wait()

	// If any failed, return aggregated error (but still hand back per-ext results).
	var failed int
	for _, r := range results {
		if r.Err != nil {
			failed++
		}
	}
	if failed == 0 && opts.Reload {
		if err := fwconsoleReload(ctx, opts.FwconsolePath, opts.FwconsoleArgs, opts.ReloadTimeout, opts.Logger); err != nil {
			return results, fmt.Errorf("fwconsole reload: %w", err)
		}
	}
	if failed > 0 {
		return results, fmt.Errorf("%d/%d deletions failed", failed, len(results))
	}
	return results, nil
}

// DeleteOne convenience helper for a single extension.
func DeleteOne(ctx context.Context, ext string, opts Options) (Result, error) {
	res, err := Delete(ctx, []string{ext}, opts)
	if len(res) == 1 {
		return res[0], err
	}
	return Result{}, err
}

// --- internals ---

func writeHelper(tempDir, confPath string) (path string, cleanup func(), err error) {
	content := strings.ReplaceAll(phpHelper, "%FREEPBX_CONF%", confPath)
	f, err := os.CreateTemp(tempDir, "fpbx-extdel-*.php")
	if err != nil {
		return "", func() {}, err
	}
	name := f.Name()
	if _, err := f.WriteString(content); err != nil {
		_ = f.Close()
		_ = os.Remove(name)
		return "", func() {}, err
	}
	if err := f.Close(); err != nil {
		_ = os.Remove(name)
		return "", func() {}, err
	}
	cleanup = func() { _ = os.Remove(name) }
	return name, cleanup, nil
}

func runPHP(ctx context.Context, phpPath, helperPath, ext string) (string, error) {
	cmd := exec.CommandContext(ctx, phpPath, helperPath, ext)
	out, err := cmd.CombinedOutput()
	return string(out), err
}
/*
func fwconsoleReload(ctx context.Context) error {
	bin := "fwconsole"
	if runtime.GOOS == "windows" {
		// Unlikely on PBX, but just in case:
		bin = "fwconsole.exe"
	}
	cmd := exec.CommandContext(ctx, bin, "reload")
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("%v: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}
*/
// --- replace your fwconsoleReload with this version ---
func fwconsoleReload(parent context.Context, bin string, extra []string, to time.Duration, lg *log.Logger) error {
	if bin == "" {
		bin = "fwconsole"
	}
	args := append(append([]string{}, extra...), "reload")

	ctx := parent
	var cancel context.CancelFunc
	if to > 0 {
		ctx, cancel = context.WithTimeout(parent, to)
		defer cancel()
	}

	cmd := exec.CommandContext(ctx, bin, args...)
	out, err := cmd.CombinedOutput()

	if lg != nil {
		lg.Printf("→ Running: %s %s", bin, strings.Join(args, " "))
		if len(out) > 0 {
			lg.Printf("fwconsole output:\n%s", strings.TrimSpace(string(out)))
		}
	}

	if err != nil {
		if errors.Is(ctx.Err(), context.DeadlineExceeded) {
			return fmt.Errorf("fwconsole reload timed out after %s", to)
		}
		return fmt.Errorf("%v: %s", err, strings.TrimSpace(string(out)))
	}
	return nil
}


// Helper to join for logs / errors.
func join(a []string) string {
	return strings.Join(a, ",")
}

// Sanity: make absolute path (useful in logs).
func abs(p string) string {
	ap, _ := filepath.Abs(p)
	return ap
}
