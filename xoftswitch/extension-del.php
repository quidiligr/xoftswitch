#!/usr/bin/env php
<?php
if (!isset($argv[1])) { echo "Usage: extension-del.php <ext>\n"; exit(1); }
include '/etc/freepbx.conf';
$FreePBX = FreePBX::Create();
$device = $FreePBX->Core->getDevice($argv[1]);
$user   = $FreePBX->Core->getUser($argv[1]);
if ($device["user"]) {
  $FreePBX->Core->delDevice($argv[1]);
  $FreePBX->Core->delUser($device["user"]);
} elseif ($user) {
  $FreePBX->Core->delUser($argv[1]);
} else {
  echo "Not found\n";
}