# GitHub Repository Setup

Your local git repository is initialized and ready. Follow these steps to create the GitHub repository and push your code:

## Option 1: Create Repository on GitHub Web Interface

1. Go to https://github.com/new
2. Repository name: `guitar-training`
3. Description: `Guitar training application built with Go`
4. Choose **Public** or **Private**
5. **DO NOT** initialize with README, .gitignore, or license (we already have these)
6. Click **Create repository**

7. After creating, run these commands:

```bash
git remote add origin https://github.com/YOUR_USERNAME/guitar-training.git
git branch -M main
git push -u origin main
```

Replace `YOUR_USERNAME` with your GitHub username.

## Option 2: Install GitHub CLI and Create Automatically

If you have GitHub CLI installed and authenticated:

```bash
gh repo create guitar-training --public --description "Guitar training application built with Go" --source=. --push
```

To install GitHub CLI:
- macOS: `brew install gh`
- Linux: See https://github.com/cli/cli/blob/trunk/docs/install_linux.md
- Windows: See https://github.com/cli/cli/blob/trunk/docs/install_windows.md

After installation, authenticate:
```bash
gh auth login
```

## Option 3: Use the Helper Script

A helper script (`.github-setup.sh`) is included. First, create the repository on GitHub (Option 1, steps 1-6), then:

1. Edit `.github-setup.sh` and update `GITHUB_USER` with your GitHub username
2. Run: `./.github-setup.sh`

## Verify Setup

After pushing, verify everything worked:

```bash
git remote -v
git log --oneline
```

Your repository should now be available at:
`https://github.com/YOUR_USERNAME/guitar-training`
