# Encrypt a PDF
```sh
qpdf --encrypt password password 128 --use-aes=y -- ./input.pdf ./input.protected.pdf
```