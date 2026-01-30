# vault

![vault banner](public/banner.png)

A lightweight CLI tool for managing reusable React and Next.js components.

Save components from any project into a local registry, list them, and sync them with a GitHub repository for reuse across multiple codebases. Instead of copy-pasting components between projects or maintaining separate component libraries, vault provides a fast, filesystem-based workflow that fits naturally into existing development habits.

## Features

- **Save** - Store `.tsx` and `.jsx` components in a local vault organized by file type
- **List** - View all saved components at a glance, grouped by TSX and JSX
- **Global Storage** - Components are stored in `~/.vault/` (user's home directory) for access from any project
- **Duplicate Detection** - Case-insensitive duplicate checking prevents accidental overwrites
- **Sync** - Push/pull components to GitHub for cross-project sharing _(coming soon)_

## Installation

### Using Go Install (Recommended)

```bash
go install github.com/Ayroid/vault@latest
```

Make sure `$GOPATH/bin` (usually `~/go/bin`) is in your `PATH`:

```bash
# Add to your ~/.bashrc or ~/.zshrc
export PATH="$PATH:$(go env GOPATH)/bin"
```

After installation, you can save files from any project:

```bash
vault save ./components/Button.tsx
```

### From Source

```bash
git clone https://github.com/Ayroid/vault.git
cd vault
go build -o vault
```

### Move to PATH (optional, if building from source)

```bash
# Linux/macOS
sudo mv vault /usr/local/bin/

# Or add to your local bin
mv vault ~/.local/bin/
```

## Usage

### Save a component

```bash
vault save <path-to-file>
```

Save `.tsx` or `.jsx` component files to your local vault:

```bash
vault save ./components/Button.tsx
vault save ./components/Card.jsx
```

Save with a custom name:

```bash
vault save ./Button.tsx --name RoundButton.tsx
```

**Note:** Only `.tsx` and `.jsx` files are supported. Components are automatically organized into separate folders based on their extension.

### List saved components

```bash
vault list
```

View all components currently stored in your vault, organized by file type:

```
TSX Components
├── Button.tsx
├── Card.tsx
└── Modal.tsx
JSX Components
├── Header.jsx
└── Footer.jsx
```

## How It Works

vault uses a simple filesystem-based approach:

1. Components are stored globally in `~/.vault/components/` (user's home directory)
2. TSX components are saved to `~/.vault/components/tsx/`
3. JSX components are saved to `~/.vault/components/jsx/`
4. Files are saved with their original filenames (or custom names via `--name`)
5. Duplicate detection is case-insensitive to prevent naming conflicts
6. No configuration files or databases required

## Storage Location

Components are stored in the user's home directory:

```
~/.vault/                   # e.g., /home/username/.vault/
└── components/
    ├── tsx/                # TypeScript React components
    │   ├── Button.tsx
    │   └── Modal.tsx
    └── jsx/                # JavaScript React components
        ├── Header.jsx
        └── Footer.jsx
```

## Commands Reference

| Command | Description |
|---------|-------------|
| `vault save <file>` | Save a component to the vault |
| `vault save <file> --name <name>` | Save a component with a custom name |
| `vault list` | List all saved components |

## Roadmap

- [x] Custom naming with `--name` flag
- [x] Separate TSX and JSX folders
- [x] Global storage in user's home directory
- [x] Case-insensitive duplicate detection
- [ ] GitHub sync support
- [ ] Component categories/tags
- [ ] Search functionality
- [ ] Config file support
- [ ] `vault get` command to copy components into projects

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

MIT
