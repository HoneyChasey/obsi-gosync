# obsi-gosync
Project write in golang, backup encrypted in cloud or in your homelab

## What file is used for ? 
logique.go = call api and code cannot be imported 

cmd/root = racine command (ogs)

main.go entry point 


## Explication 

To build a zip archive you need the metadata of each file. To do this we call file.Stat() to get an os.FileInfo, then we use zip.FileInfoHeader to create a header from it, and finally zipWriter.CreateHeader to write it into the archive.

# TODO 
- Maybe use cryptomator cli to create crypted vault and encrypt directly to the pc. Thx to cryptomator, unecesserarry to uncrypt and crypt vault for ich modification.
