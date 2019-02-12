const path = require("path")

module.exports = (baseConfig, env, defaultConfig) => {
  defaultConfig.module.rules.push({
    test: /\.(ts|tsx)$/,
    include: path.resolve(__dirname, '../'),
    loader: require.resolve("ts-loader")
  });

  // addon-storysource使うときだけ
  defaultConfig.module.rules.push({
    test: /\.stories\.jsx?$/,
    loaders: [
      {
        loader: require.resolve("@storybook/addon-storysource/loader"),
        options: { parser: "typescript" }
      }
    ],
    enforce: "pre"
  });
  defaultConfig.resolve.extensions.push(".ts", ".tsx")

  return defaultConfig
}
