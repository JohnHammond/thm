# Gotta Catch Them All!

> John Hammond | August 13th, 2020

----------------------------------

# Task 1. Can You Catch'em All? 

> Remember to connect to the VPN network using OpenVPN, It may take some time for the machine to properly deploy.
> 
> You can also deploy your own Kali Linux machine, and control it in your browser using the provided Kali machine (Subscription Required).
> 
> Enjoy the room!



## 1. Find the Grass-Type Pokemon


Looking at the website you can find some odd HTML tags `pokemon` and `hack_the_pokemon` which
will be SSH credentials to log in.

Found in the `P0k3m0n.zip` file in the `pokemon` user Desktop.

```
PoKeMoN{Bulbasaur}
```

While looking at files with `ls -lR`, I found a `.cplusplus` file nested that includes the password for the `ash` account.

```
ash:pikapika
```

This `ash` account can run any command as `root` via `sudo`, so you can own the box.

## 2. Find the Water-Type Pokemon

Find in the `/var/www/html` directory.

```
Squirtle_SqUaD{Squirtle}
```

## 3. Find the Fire-Type Pokemon

Find in `/etc`....

I just found this with the syntax:

```
find -name fire-type.txt
```

```
P0k3m0n{Charmander}
```

## 4. Who is Root's Favorite Pokemon?

Found in `/home`.

