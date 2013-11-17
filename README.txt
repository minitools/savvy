Savvy
=====
Incredibly basic backup utility that:
1. scans local folders, extracting last modification date
2. archives folders that were changed since last scan

Folders are archived into a fixed destination folder (currently hard-coded)
as Zip files with the following name:

{folder name}-{dateYYMMDD-HHMMSS}.zip


Version history
---------------
0.2		Ignore certain directories
		- If a file named "savvyignore" is stored into the top level of a directory, that directory will not be scheduled for backup

0.1		Initial version


TODO:
- Default to local destination directory

- Add way to specify destination directory
  1. first time, by passing an argument to the program
     the program saves the destination into a .savvy file inside the current directory
  2. if the program starts and finds the .savvy file inside current directory, 
     reads it to extract destination directory
  3. if it doesn't find a .savvy directory and no destination is specified, it logs an error and exits

- Add test suite

