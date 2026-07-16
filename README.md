# photopad

A picture's worth a thousand pads.™

Photopad is a small CLI tool which will take plain text and a jpeg image as input, convert the image's pixels into a "[pre-shared key](https://en.wikipedia.org/wiki/Pre-shared_key)", and then use this key as a cryptographic pad to encrypt the provided text, and vice versa.

> [!WARNING]
> This project was intended as a learning exercise for the author. It is not intended to be a secure form of encryption, and I strongly advise anyone that might be considering using it for purposes other than low stakes tomfoolery to pause and research the many better alternatives that were created by smarter people with more intention.

# Installation

The only way to use this tool is to clone the repo and compile the code locally. This project is written in [Go](https://go.dev). Once you have Go installed, you can run `go run photopad-cli.go`. This tool is not distributed or published anywhere else on the web. If you see it somewhere in the wild, you should not assume the author of this respository endorses or condones that use.

# Usage

The API for this tool is limited to two commands. You are either encrypting plaintext or decrypting ciphertext. At this stage in the project's maturity, only `.txt` files are supported for both input and output. Similarly, although there are many file formats that exist in the world, this silly little program is only able to convert `.jp(e)g` files to the cryptographic pad necessary to complete the operation. All other file formats will result in an error.

## Encrypt

Convert a plaintext file to an encrypted ciphertext file based on the pixel values of a .jpg file. 

```nix
> ./photopad-cli encrypt --input ./path/to/plain.txt \
  --key ./path/to/photo.jpg [--output ./custom/path/for/cipher.txt]
```

## Decrypt

Convert a ciphertext file to a decrypted plaintext file based on the pixel values of a .jpg file.

```nix
> ./photopad-cli encrypt --input ./path/to/cipher.txt \
  --key ./path/to/photo.jpg [--output ./custom/output/plain.txt]
```

# Still Have Questions?

This `README` intentionally doesn't dive too deeply on the inner operations. If you made it to the bottom and you want to learn more about the internals (without having to read my amature go code), you can read the [blog post]() I wrote when I finished the first version.
