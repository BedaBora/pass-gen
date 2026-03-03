# 🛡️ PassGen
A Random Password Generator

<b>A cryptographically secure command-line password generator built in Go.</b>

PassGen uses the `crypto/rand` library to ensure that every password generated is unpredictable and safe for sensitive accounts. Unlike tools that use standard math-based randomness, PassGen is suitable for security-critical applications.

# ✨ Features

- Secure: Uses OS-level entropy for true randomness.
- Customizable: Control length, digits, and special characters.
- Cross-Platform: Single-binary execution for Windows, Linux, and macOS.
- Fast: Zero dependencies and lightning-fast execution.


# 🚀 Installation
<b>From Source</b>

Ensure you have [Go](https://go.dev/) installed, then:
```sh
git clone https://github.com/BedaBora/pass-gen.git
cd passgen
go build -o passgen
```

<b>Using the Makefile</b>

If you have `make` installed, you can build for all platforms at once:
```sh
make all
```

The binaries will be located in the /bin directory.

# 🛠️ Usage

Run the executable followed by the desired flags.

<b>Basic Command</b>
```sh
./passgen -l 24
```

<b>Options</b>
| Flag | Description | Default |
|--- | --- | --- |
| -l |	Length of the password |	16 |
| -d |	Include numbers (true/false) |	true |
| -s |	Include special characters (true/false) |	true |

## Examples

Generate a 32-character complex password:
```sh
/passgen -l 32 -d=true -s=true
```

Show help message:
```sh
./passgen -h
```