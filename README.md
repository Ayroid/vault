# gocopy

A lightweight CLI tool for managing reusable React and Next.js components.

Save components from any project into a local registry, list them, and sync them with a GitHub repository for reuse across multiple codebases. Instead of copy-pasting components between projects or maintaining separate component libraries, gocopy provides a fast, filesystem-based workflow that fits naturally into existing development habits.

## Features

- **Save** - Store components in a local vault for quick access
- **List** - View all saved components at a glance
- **Sync** - Push/pull components to GitHub for cross-project sharing *(coming soon)*

## Installation

### From Source

```bash
git clone https://github.com/yourusername/gocopy.git
cd gocopy
go build -o gocopy
```

### Move to PATH (optional)

```bash
# Linux/macOS
sudo mv gocopy /usr/local/bin/

# Or add to your local bin
mv gocopy ~/.local/bin/
```

## Usage

### Save a component

```bash
gocopy save <path-to-file>
```

Save any file to your local vault:

```bash
gocopy save ./components/Button.tsx
gocopy save ./hooks/useAuth.ts
gocopy save ./lib/utils.ts
```

### List saved components

```bash
gocopy list
```

View all components currently stored in your vault.

## How It Works

gocopy uses a simple filesystem-based approach:

1. Components are stored in `.vault/components/` relative to where you run the command
2. Files are saved with their original filenames
3. No configuration files or databases required

## Project Structure

```
your-project/
├── .vault/
│   └── components/     # Your saved components live here
├── src/
└── ...
```

## Roadmap

- [ ] Custom naming with `--name` flag
- [ ] GitHub sync support
- [ ] Component categories/tags
- [ ] Search functionality
- [ ] Config file support

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

## License

MIT
