# gore-utils

An unnecessary, unapologetic reimplementation of already battle-tested software that absolutely nobody needed.

## Warning

This project has no warranty, no actual roadmap, and no regrets. It exists because curiosity won, and curiosity has poor risk assessment skills.

This software is provided "as is", "as was" and possibly "as should not have been", so use it at your own risk. Or don't.

## Manifesto

> I wonder if I can
>
> -- me, before `git init`

`gore-utils` is a slightly unhinged educational experiment written in Go (until I get bored of the language or develop emotional attachment to it) to understand how things work under the hood, and to break them respectfully.

We believe in:

- Rewriting things that already work
- Learning by rebuilding
- Accepting that performance regressions build character
- Shipping anyway, because that's what programmers do

We do not aim to replace coreutils or other programs, we aim to understand them and, possibly, reimagine them.

If you are looking for production-ready software, this is not it. If you are looking for engligthment through mild suffering, welcome.

## Zen

May thy pipes never break.

May thy flags always parse.

May thy panics be educational.

## Roadmap (or maybe Wishlist)

> NOTE: names are WIP

In no particular order and with no guarantee that I will actually implement them:

- `copy` high-performance byte-shoveler
- `list` fancy directory lister (maybe with tree support if I feel brave)
- `spill` cat with a different name, maybe include head and tail with flags, less-style pagination would be cool
- `dump` like hexdump but with a default -C
- `move` move and rename files, with interactive mode to prevent data loss
- `tap` create files or update timestamp of a file (maybe dirs too with flag -d?)
- `del` rm but safer, trust me
- `count` wc but without the toilet reference
- `loc` find without the unnecessary complexity
- `seek` search by text patterns
- `sub` replace by text patterns, highly dangerous
- `mory` in-memory database
- `comp` file compressor (Huffman, unless I get skill issues mid-implementation)
- `net` network related tools, maybe a -p for ping and basic requests functionality
- `pass` cli based password manager
- and probably chmod, realpath, base64, cksum, torrent

If I ever rewrite something in C, get ready for `chore-utils`
