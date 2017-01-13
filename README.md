# socialcp

# Install

```
make setup
make build  
make install
```

# Usage

In order to use `socialcp` you must first register your client in the server. This is done by the command `register` passing your e-mail as a parameter:

`socialcp register your@email.com`

The e-mail is your personal identifier in the server.

You may add your friends with the command `add` giving a slot number and an e-mail identifier:

`socialcp add slot# friend@email.com`

`slot#` is any number from 0-9.

Finally, you must register the hotkeys for the `send` and `recv` commands, which is SO dependant. Both commands are explained below.

The `send` command sends the content of your clipboard to the remote clipboard slot identified by the key pair composed by your identifier and your friend identifier. The `recv` works the same way, but receiving (copying) content from the remote clipboard slot to your local clipboard. Please note that there is one slot for each direction.

The diagram below shows the process of copying something on user A's computer and pasting it on user B's computer:
```
                                        remote clipboard (server)
                                      (one slot for each direction)
local clipboard                    /--------------------------------\                      local clipboard
   (user A)                        | origin | destination | content |                         (user B)
 /-----------\   socialcp send #   |--------------------------------|   socialcp recv #   /-------------\
 |   hello   | ==================> | user A | user B      | hello   | ==================> |    hello    |
 \-----------/                     |--------------------------------|                     \-------------/
      /\                           | user B | user A      | world   |                           ||
      ||                           |--------------------------------|                           ||  paste
      ||  copy                     :                                :                           || (Cmd+V)
      || (Cmd+C)                   :   slots for all other users    :                           ||
      ||                           :                                :                           \/
                                   |--------------------------------|                           
    hello                          | user n | user m      | xpto    |                         hello
                                   \--------------------------------/
```

Basically, in order to send a _copy_ of something to your friend you must first copy it yourself with the usual copy hotkey (Cmd+C for OSX or Ctrl+C for Linux/Windows) and then issue a `send` command passing as a parameter the slot number for a previously configured user, i.e., `socialcp send 1`. This command can be issued at the command line but ideally one would associate it with a hotkey to run the command line itself.

The process to receive a _paste_ works similarly, but you must first run a `recv` command to receive the data from the remote slot and then issue a paste command with the usual hotkey (Cmd+V for OSX and Ctrl+V for Linux/Windows).

# Examples

1. Register your user: `socialcp register me@somewhere.com`
    
2. Configure your friends: `socialcp add 1 myfriend@anywhere.com`  

3. Map the shortcuts (SO dependant)
  - Command line to send to remote clipboard: `socialcp send 1`  
  - Command line to receive from remote clipboard: `socialcp recv 1` 

4. Enjoy :)
