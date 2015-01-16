Savvy [![Build Status](https://travis-ci.org/minitools/savvy.svg?branch=master)](https://travis-ci.org/minitools/savvy)
=====
Incredibly basic backup utility that:
1. scans local folders, extracting last modification date
2. archives folders that were changed since last scan

Folders are archived into a configurable destination folder
as Zip files with the following name:

{folder name}-{dateYYMMDD-HHMMSS}.zip


Supported platforms
-------------------
- Unix/Linux
- MacOS
- Windows


Version history
---------------
0.4.1	Fixed issue with destination path being overwritten.
		Also, checks existence of destination before proceeding with backup.

0.4		Added options:
			-n 		don't perform backup (no op)
			-v 		verbose output

0.3		Destination folder can be specified from command line:
			-dest <directory>

		This setting is saved in a JSON-encoded configuration file,
			$HOME/.savvy.config

0.2		Ignore certain directories
		- If a file named "savvyignore" is stored into the top level of a directory, that directory will not be scheduled for backup

0.1		Initial version


To Do
-----
- When running with -n, show exactly which directories would be backed-up

- Global configuration file, with support for
  - exclude [pattern]
  - recurse [pattern]

- Specify source directory as argument

- Destination folder is excluded automatically

- Add test suite

Questions, ideas?
-----------------
Contact me at Tom Paoletti <tpaoletti_JUSTDROP@users.sf.net>
