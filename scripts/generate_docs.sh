# Generate go doc for all packages in pkg folder
for d in pkg/*; do
  # Check if the entry is a directory
  if [ -d "${d}" ]; then
    # Generate the documentation for the package
    (
      cd "${d}"

      # display documentation
      go doc -all
    )
  fi
done
