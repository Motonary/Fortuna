module.exports = {
  setupTestFrameworkScriptFile: "./node_modules/jest-enzyme/lib/index.js",
  collectCoverage: true,
  transform: {
    "^.+\\.(ts|tsx)$": "ts-jest",
    "\\.(jpg|jpeg|png|gif|eot|otf|webp|svg|ttf|woff|woff2|mp4|webm|wav|mp3|m4a|aac|oga)$": "<rootDir>src/__mocks__/fileTransformer.js"
  },
  testMatch: [
    "**/*.test.(ts|tsx)"
  ],
  testPathIgnorePatterns: [
    "/node_modules/",
    "!/node_modules/types"
  ],
  moduleFileExtensions: [
    "ts",
    "tsx",
    "js",
    "jsx",
    "json"
  ],
  moduleNameMapper: {
    "\\.(css)$": "identity-obj-proxy"
  },
  verbose: true
}
