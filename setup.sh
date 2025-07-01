#!/bin/bash

# 1. Create a new directory for the project
mkdir -p {{.RepoName}}

# 2. Copy the template files to the new directory
cp -r _template/* {{.RepoName}}/

# 3. Replace the placeholders in the template files with the actual values
sed -i '' "s/{{.RepoName}}/{{.RepoName}}/g" {{.RepoName}}/README.md
sed -i '' "s/{{.ModuleName}}/{{.ModuleName}}/g" {{.RepoName}}/go.mod