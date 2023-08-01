
# GoRat
GoRat is a simple Remote Access Trojan (RAT) built in GoLang, GoRat allows connections from machines and allows the hosting machine to do certain actions on the machines, GoRat uses its own API instead of Websockets as its intended to be hosted on a Windows RDP, GoRat was a basic project to support me while learning GoLang and is far from perfect, this means it will have bugs and is not the best for production use.

## GoRat GUI:
GoRat has its very own GUI made using `giu`, the GUI automatically refreshes to show new data and is quite speedy. The GUI allows easier communication between the client and the server and makes the RAT just better overall.

![App Screenshot](https://raw.githubusercontent.com/NotKatsu/GoRat/main/assets/example.png)


## Running GoRat
Although the project is not done right now it can still be ran however this will only run the server (API and GUI) because the Client is currently not done however you can use `Insomnia` to replicate the API requests or code your own Client.

Although the Client is not done and the Server isn't fully done you can still implement "events" easily to the API and GUI with little to no prior experience in GoLang.


# Warning
This project was intended to be used for machines you own and not other people machines hence why it uses an API over Websockets, although i can't really stop people from doing this i really encourage people not to and by modifying the program or adding/removing anything that makes it now your software and i take no accountablility for what is done with it.