#!/bin/bash
# Helper script to push to GitHub
# This script assumes you've created the repository on GitHub first

REPO_NAME="guitar-training"
GITHUB_USER="paulgreig"  # Update this with your GitHub username

echo "Setting up GitHub remote..."
git remote add origin "https://github.com/${GITHUB_USER}/${REPO_NAME}.git"

echo "Renaming branch to main..."
git branch -M main

echo "Pushing to GitHub..."
git push -u origin main

echo "Done! Repository should be available at: https://github.com/${GITHUB_USER}/${REPO_NAME}"
