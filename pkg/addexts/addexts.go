package addexts

import (
	"context"
	cryptoRand "crypto/rand"
	"encoding/csv"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

type OptionsRange struct {
	TemplatePath string
	OutPath      string

	Start, End   int
	NamePattern  string // e.g. "WebRTC {ext}"
	//Name  string
	EmailPattern string // e.g. "xo{ext}@xoftphone.com"
	//Email string
	Secret string
	MediaEncryption string //x5
	DirectMedia string //x7
	WebRtc string //x6
	
	//SecretBytes  int    // bytes -> hex-encoded length is 2*bytes

	// Optional side-effects:
	FwconsolePath string        // default "fwconsole"
	DoImport      bool          // run `fwconsole bulkimport`
	DoReload      bool          // run `fwconsole reload` after import
	Timeout       time.Duration // per fwconsole command

	// Testability / DI:
	RandReader io.Reader // default crypto/rand.Reader
	Logger     *log.Logger
}

type Options struct {
	TemplatePath string
	OutPath      string

	Extension   string
	NamePattern  string // e.g. "WebRTC {ext}"
	//Name  string
	EmailPattern string // e.g. "xo{ext}@xoftphone.com"
	//Email string
	Secret string
	MediaEncryption string //x5
	DirectMedia string //x7
	WebRtc string //x6
	//SecretBytes  int    // bytes -> hex-encoded length is 2*bytes

	// Optional side-effects:
	FwconsolePath string        // default "fwconsole"
	DoImport      bool          // run `fwconsole bulkimport`
	DoReload      bool          // run `fwconsole reload` after import
	Timeout       time.Duration // per fwconsole command

	// Testability / DI:
	RandReader io.Reader // default crypto/rand.Reader
	Logger     *log.Logger
}

type GeneratedExt struct {
    //ExtNumber   int
	Extension   string
    DisplayName string
    Email       string
    Secret      string

	MediaEncryption string //x5
	DirectMedia string //x7
	WebRtc string //x6
}

//var FwconsolePath = "fwconsole"
//var DoImport = true
//var DoReload = true

func (o *OptionsRange) setDefaults() {
	
	/*
	if o.SecretBytes <= 0 {
		o.SecretBytes = 16
	}
		*/
	//o.DoReload = true
	//o.DoImport = true

	if o.FwconsolePath == "" {
		o.FwconsolePath = "fwconsole"
	}
	
	if o.Timeout == 0 {
		o.Timeout = 10 * time.Minute
	}
	if o.RandReader == nil {
		o.RandReader = cryptoRand.Reader
	}
	if o.Logger == nil {
		o.Logger = log.New(os.Stdout, "", log.LstdFlags)
	}
}

func (o *Options) setDefaults() {
	
	/*
	if o.SecretBytes <= 0 {
		o.SecretBytes = 16
	}
		*/
		/*
	if o.FwconsolePath == "" {
		o.FwconsolePath = "fwconsole"
	}
	*/
	//o.DoReload = true
	//o.DoImport = true

	if o.FwconsolePath == "" {
		o.FwconsolePath = "fwconsole"
	}
	

	if o.Timeout == 0 {
		o.Timeout = 10 * time.Minute
	}
	if o.RandReader == nil {
		o.RandReader = cryptoRand.Reader
	}
	if o.Logger == nil {
		o.Logger = log.New(os.Stdout, "", log.LstdFlags)
	}
}



// Add this type somewhere near Options:


// New signature returns the generated list.
func GenerateRange(ctx context.Context, opts OptionsRange) ([]GeneratedExt, error) {
    opts.setDefaults()
	//if err != nil {
		//fmt.Printf("marshal opts: %v\n", err)
	//} else {
	//	fmt.Println("opts:", string(b))
	//}
	opts.Logger.Printf("opts: %+v", opts)

    if opts.End < opts.Start {
        return nil, fmt.Errorf("invalid range: end < start (%d < %d)", opts.End, opts.Start)
    }
    if opts.TemplatePath == "" {
        return nil, fmt.Errorf("TemplatePath is required")
    }
    if opts.OutPath == "" {
        return nil, fmt.Errorf("OutPath is required")
    }

    // read template
    tf, err := os.Open(opts.TemplatePath)
    if err != nil {
        return nil, fmt.Errorf("open template: %w", err)
    }
    defer tf.Close()

    r := csv.NewReader(tf)
    r.FieldsPerRecord = -1

    header, err := r.Read()
    if err != nil {
        return nil, fmt.Errorf("read header: %w", err)
    }
    if len(header) == 0 {
        return nil, fmt.Errorf("template header is empty")
    }
    baseRow, err := r.Read()
    if err == io.EOF {
        return nil, fmt.Errorf("template must include base row (row 2)")
    }
    if err != nil {
        return nil, fmt.Errorf("read base row: %w", err)
    }

    if len(baseRow) < len(header) {
        tmp := make([]string, len(header))
        copy(tmp, baseRow)
        baseRow = tmp
    } else if len(baseRow) > len(header) {
        baseRow = baseRow[:len(header)]
    }

    // ensure out dir
    if dir := filepath.Dir(opts.OutPath); dir != "" && dir != "." {
        if err := os.MkdirAll(dir, 0o755); err != nil {
            return nil, fmt.Errorf("mkdir %s: %w", dir, err)
        }
    }

    // write output
    out, err := os.Create(opts.OutPath)
    if err != nil {
        return nil, fmt.Errorf("create out: %w", err)
    }
    defer out.Close()

    w := csv.NewWriter(out)

    if err := w.Write(header); err != nil {
        return nil, fmt.Errorf("write header: %w", err)
    }

    // Collect results to return
    results := make([]GeneratedExt, 0, opts.End-opts.Start+1)

    for ext := opts.Start; ext <= opts.End; ext++ {
        x1 := fmt.Sprintf("%d", ext)                                   // {ext}
        x2 := strings.ReplaceAll(opts.NamePattern, "{ext}", x1)        // display name
        secret := opts.Secret
        if secret == "" {
            newSecret, err := generateSecret(opts.RandReader, 16)
            if err != nil {
                return nil, fmt.Errorf("generate secret for %d: %w", ext, err)
            }
            secret = newSecret
        }
        x4 := strings.ReplaceAll(opts.EmailPattern, "{ext}", x1)       // email
		x5 := opts.MediaEncryption
		x6 := opts.WebRtc   
		x7 := opts.DirectMedia   

        record := make([]string, len(header))
        for i, cell := range baseRow {
            record[i] = replaceTokens(cell, x1, x2, secret, x4,x5,x6,x7)
        }
        if err := w.Write(record); err != nil {
            return nil, fmt.Errorf("write row for %d: %w", ext, err)
        }

        results = append(results, GeneratedExt{
            Extension:   x1,
            DisplayName: x2,
            Email:       x4,
            Secret:      secret,
			WebRtc: x6,
			MediaEncryption: x5,
			DirectMedia: x7,
        })
    }

    w.Flush()
    if err := w.Error(); err != nil {
        return nil, fmt.Errorf("flush: %w", err)
    }

    opts.Logger.Printf("Wrote %s for extensions %d..%d using %s",
        opts.OutPath, opts.Start, opts.End, opts.TemplatePath)

		
    if opts.DoImport {
        if err := runWithTimeout(ctx, opts.FwconsolePath, []string{"bulkimport", "--type=extensions", opts.OutPath}, filepath.Dir(opts.OutPath), opts.Timeout, opts.Logger); err != nil {
            return nil, fmt.Errorf("fwconsole bulkimport: %w", err)
        }
        opts.Logger.Println("fwconsole bulkimport completed successfully.")
        if opts.DoReload {
            if err := runWithTimeout(ctx, opts.FwconsolePath, []string{"reload"}, filepath.Dir(opts.OutPath), opts.Timeout, opts.Logger); err != nil {
                return nil, fmt.Errorf("fwconsole reload: %w", err)
            }
            opts.Logger.Println("fwconsole reload completed successfully.")
        }
    } else {
        opts.Logger.Println("Skipping fwconsole bulkimport (DoImport=false).")
    }

    return results, nil
}

func Generate(ctx context.Context, opts Options) (*GeneratedExt, error) {
    opts.setDefaults()

    // Validate required options
    if strings.TrimSpace(opts.TemplatePath) == "" {
        return nil, fmt.Errorf("TemplatePath is required")
    }
    if strings.TrimSpace(opts.OutPath) == "" {
        return nil, fmt.Errorf("OutPath is required")
    }

    ext := strings.TrimSpace(opts.Extension)
    if ext == "" {
        return nil, fmt.Errorf("Extension is required")
    }

    // Open and parse template CSV
    tf, err := os.Open(opts.TemplatePath)
    if err != nil {
        return nil, fmt.Errorf("open template: %w", err)
    }
    defer tf.Close()

    r := csv.NewReader(tf)
    r.FieldsPerRecord = -1

    header, err := r.Read()
    if err != nil {
        return nil, fmt.Errorf("read header: %w", err)
    }
    if len(header) == 0 {
        return nil, fmt.Errorf("template header is empty")
    }

    baseRow, err := r.Read()
    if err == io.EOF {
        return nil, fmt.Errorf("template must include base row (row 2)")
    }
    if err != nil {
        return nil, fmt.Errorf("read base row: %w", err)
    }

    // Normalize baseRow length to header length
    if len(baseRow) < len(header) {
        tmp := make([]string, len(header))
        copy(tmp, baseRow)
        baseRow = tmp
    } else if len(baseRow) > len(header) {
        baseRow = baseRow[:len(header)]
    }

    // Ensure output directory exists
    if dir := filepath.Dir(opts.OutPath); dir != "" && dir != "." {
        if err := os.MkdirAll(dir, 0o755); err != nil {
            return nil, fmt.Errorf("mkdir %s: %w", dir, err)
        }
    }

    // Create output file
    out, err := os.Create(opts.OutPath)
    if err != nil {
        return nil, fmt.Errorf("create out: %w", err)
    }
    defer out.Close()

    w := csv.NewWriter(out)

    if err := w.Write(header); err != nil {
        return nil, fmt.Errorf("write header: %w", err)
    }

    // Prepare tokens
    x1 := ext // {ext}

    // Display name
    x2 := strings.ReplaceAll(opts.NamePattern, "{ext}", x1)

    // Secret
    secret := opts.Secret
    if secret == "" {
        newSecret, err := generateSecret(opts.RandReader, 16)
        if err != nil {
            return nil, fmt.Errorf("generate secret for %s: %w", ext, err)
        }
        secret = newSecret
    }

    // Email
    x4 := strings.ReplaceAll(opts.EmailPattern, "{ext}", x1)

    // Media/WebRTC/DirectMedia (coerce to string if your replaceTokens expects strings)
    x5 := opts.MediaEncryption
    x6 := opts.WebRtc        // if bool: strconv.FormatBool(opts.WebRtc)
    x7 := opts.DirectMedia   // if bool: strconv.FormatBool(opts.DirectMedia)

    // Build record
    record := make([]string, len(header))
    for i := 0; i < len(header); i++ {
        cell := baseRow[i]
        record[i] = replaceTokens(cell, x1, x2, secret, x4, x5, x6, x7)
    }

    if err := w.Write(record); err != nil {
        return nil, fmt.Errorf("write row for %s: %w", ext, err)
    }

    w.Flush()
    if err := w.Error(); err != nil {
        return nil, fmt.Errorf("flush: %w", err)
    }

    if opts.Logger != nil {
        opts.Logger.Printf("Wrote %s for extension %s using %s", opts.OutPath, ext, opts.TemplatePath)
    }

    // Optional import + reload via fwconsole
    if opts.DoImport {
        if err := runWithTimeout(
            ctx,
            opts.FwconsolePath,
            []string{"bulkimport", "--type=extensions", opts.OutPath},
            filepath.Dir(opts.OutPath),
            opts.Timeout,
            opts.Logger,
        ); err != nil {
            return nil, fmt.Errorf("fwconsole bulkimport: %w", err)
        }
        if opts.Logger != nil {
            opts.Logger.Println("fwconsole bulkimport completed successfully.")
        }

        if opts.DoReload {
            if err := runWithTimeout(
                ctx,
                opts.FwconsolePath,
                []string{"reload"},
                filepath.Dir(opts.OutPath),
                opts.Timeout,
                opts.Logger,
            ); err != nil {
                return nil, fmt.Errorf("fwconsole reload: %w", err)
            }
            if opts.Logger != nil {
                opts.Logger.Println("fwconsole reload completed successfully.")
            }
        }
    } else if opts.Logger != nil {
        opts.Logger.Println("Skipping fwconsole bulkimport (DoImport=false).")
    }

    // Return pointer per signature
    result := GeneratedExt{
        Extension:       ext,
        DisplayName:     x2,
        Email:           x4,
        Secret:          secret,
        WebRtc:          opts.WebRtc,
        MediaEncryption: opts.MediaEncryption,
        DirectMedia:     opts.DirectMedia,
    }
    return &result, nil
}


// Optional: keep the old signature alive for callers that ignore the return.
func GenerateNoReturn(ctx context.Context, opts Options) error {
    _, err := Generate(ctx, opts)
    return err
}


func replaceTokens(s, x1, x2, x3, x4,x5,x6,x7 string) string {
	s = strings.ReplaceAll(s, "x1", x1) // extension
	s = strings.ReplaceAll(s, "x2", x2) // name
	s = strings.ReplaceAll(s, "x3", x3) // secret
	s = strings.ReplaceAll(s, "x4", x4) // email
	s = strings.ReplaceAll(s, "x5", x5) // media_encryption dtls/blank
	s = strings.ReplaceAll(s, "x6", x6) // webrtc yes/no
	s = strings.ReplaceAll(s, "x7", x7) // directmedia yes/no
	return s
}

func generateSecret(r io.Reader, n int) (string, error) {
	if n <= 0 {
		n = 16
	}
	b := make([]byte, n)
	// fall back to math/rand if crypto fails
	if _, err := io.ReadFull(r, b); err != nil {
		tmp := make([]byte, n)
		_, _ = rand.New(rand.NewSource(time.Now().UnixNano())).Read(tmp)
		return hex.EncodeToString(tmp), nil
	}
	return hex.EncodeToString(b), nil
}

func runWithTimeout(ctx context.Context, bin string, args []string, workDir string, timeout time.Duration, lg *log.Logger) error {
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	cmd := exec.CommandContext(ctx, bin, args...)
	cmd.Dir = workDir
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	lg.Printf("→ Running: %s %s", bin, strings.Join(args, " "))
	if err := cmd.Start(); err != nil {
		return fmt.Errorf("start %s %v: %w", bin, args, err)
	}
	if err := cmd.Wait(); err != nil {
		if ctx.Err() == context.DeadlineExceeded {
			return fmt.Errorf("timed out after %s running %s %v", timeout, bin, args)
		}
		return err
	}
	return nil
}
