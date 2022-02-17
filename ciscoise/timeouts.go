package ciscoise

import "time"

const HOTPATCH_INSTALL_TIMEOUT = time.Duration(30) * time.Minute
const HOTPATCH_ROLLBACK_TIMEOUT = time.Duration(30) * time.Minute

const PATCH_INSTALL_TIMEOUT = time.Duration(30) * time.Minute
const PATCH_ROLLBACK_TIMEOUT = time.Duration(30) * time.Minute

const HOTPATCH_INSTALL_TIMEOUT_SLEEP = time.Duration(5) * time.Minute
const HOTPATCH_ROLLBACK_TIMEOUT_SLEEP = time.Duration(3) * time.Minute
const PATCH_INSTALL_TIMEOUT_SLEEP = time.Duration(5) * time.Minute
const PATCH_ROLLBACK_TIMEOUT_SLEEP = time.Duration(3) * time.Minute
