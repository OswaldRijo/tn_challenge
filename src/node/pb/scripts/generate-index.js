const fs = require("fs");
const path = require("path");

const root = path.resolve(__dirname, "../src");

const files = [];

/**
 * Find all files in directory recursively
 */
const getFiles = (dir) => {
  fs.readdirSync(dir).forEach((file) => {
    const absolute = path.resolve(dir, file);
    const relative = path.relative(root, absolute);
    if (fs.statSync(absolute).isDirectory()) return getFiles(absolute);
    else return files.push(relative);
  });
};

getFiles(root);

// Write index.ts
const index = files
  .filter((file) => file.endsWith(".ts"))
  .map((file) => `export * from './${file.replace(/\.ts$/, "")}';`)
  .join("\n");

fs.writeFileSync(path.resolve(root, "index.ts"), index);
