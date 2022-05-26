# RWSUU_DAT2CSV

This is a Golang script that will convert a binary log file created from the Firmware to a human-readable CSV file.
The binary dat files are structured as follows:

![RWSUU DAT structure](https://github.com/remotewatersensing/RWSUU-Diagrams/blob/main/diagrams/Datastructure.png?raw=true "RWSUU DAT structure")

## How to use
Make sure you have the latest version of Go installed. Run the script in the terminal with:

> go run main.go RWSUU_00.DAT

Where RWSUU_00.DAT is an example of a filename you can use.
