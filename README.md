## zlint

This is a fork of [zmap/zlint](https://github.com/zmap/zlint) that uses the builtin packages in go.

### TLD's

In this repository is stored a list of TLD's

#### Updating the list

Pull a new version of the list from IANA, like so:

TODO(adam): Convert this into gosource so it's portable w/o the extra files needed

```
curl -o data/newgtlds.txt http://data.iana.org/TLD/tlds-alpha-by-domain.txt
```
