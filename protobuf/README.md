# Protocol Buffers Demo

The Go example demo for [Google Protocol Buffers](https://github.com/google/protobuf).

## Compile Protocol Buffers

1. Install Protocol Compiler
    [Download the package](https://github.com/google/protobuf/releases) and follow the instructions in the README
2. Install Protobuf Compiler Plugin
    ```shell
    go get -u github.com/golang/protobuf/protoc-gen-go
    ```
3. Compile Protocol Buffers
    ```shell
    protoc -I=$SRC_DIR --go_out=$DST_DIR $SRC_DIR/addressbook.proto
    ```

## Usage

### Add People

Add the first people into address-book.txt.

```shell
# cd add-people
# go build -o add-people add_people.go
# add-people address-book.txt
address-book.txt: File not found.  Creating new file.
Enter person ID number: 1
Enter name: robin
Enter email address (blank for none): test@github.com
Enter a phone number (or leave blank to finish): 12345678
Is this a mobile, home, or work phone? home
Enter a phone number (or leave blank to finish): 13187654321
Is this a mobile, home, or work phone? mobile
Enter a phone number (or leave blank to finish):

# cat address-book.txt

7
robintest@github.com"
                                                                                                                                                                                            12345678"

13187654321
```

Add the second people into address-book.txt.

```shell
$ add-people address-book.txt
Enter person ID number: 2
Enter name: eagle
Enter email address (blank for none): eagle@github.com
Enter a phone number (or leave blank to finish): 87654321
Is this a mobile, home, or work phone? work
Enter a phone number (or leave blank to finish): 13210101010
Is this a mobile, home, or work phone? mobile
Enter a phone number (or leave blank to finish):

# cat address-book.txt

7
robintest@github.com"
                                                                                                                                                                                            12345678"

13187654321
8
eagleeagle@github.com"
                                                                                                                                                                                            87654321"

13210101010
```

### List People

```shell
# cd list-people
# go build -o list-people list_people.go
# list-people address-book.txt
Person ID: 1
  Name: robin
  E-mail address: test@github.com
  Home phone #: 12345678
  Mobile phone #: 13187654321
Person ID: 2
  Name: eagle
  E-mail address: eagle@github.com
  Work phone #: 87654321
  Mobile phone #: 13210101010
```