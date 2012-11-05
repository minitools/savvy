Savvy
=====
Incredibly basic backup utility that:
1. scans local folders, extracting last modification date
2. archives folders that were changed since last scan

Folders are archived into a fixed destinamtion folder (currently hard-coded)
as Zip files with the following name:
{folder name}-{dateYYMMDD-HHMMSS}.zip
