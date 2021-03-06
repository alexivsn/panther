const path = require('path');

module.exports = {
  extends: [
    'airbnb-base', // Use Airbnb's eslinting rule for this project https://github.com/airbnb/javascript/tree/master/packages/eslint-config-airbnb-base
    'plugin:react/recommended',
    'plugin:@typescript-eslint/recommended',
    'prettier/@typescript-eslint',
    'plugin:prettier/recommended',
  ],
  env: {
    browser: true,
    node: true,
  },
  ignorePatterns: ['npm-debug.log*', 'dist', '__generated__', '*.generated.*'],
  rules: {
    'import/prefer-default-export': 0,
    'max-len': 0,
    'no-unused-vars': 'off',
    'no-underscore-dangle': 'off',
    'react/jsx-filename-extension': 0,
    'react/prop-types': 0,
    'import/no-extraneous-dependencies': 0,
    '@typescript-eslint/no-var-requires': 0,
    '@typescript-eslint/no-empty-function': 0,
    '@typescript-eslint/ban-ts-ignore': 0,
    '@typescript-eslint/explicit-function-return-type': 0,
    '@typescript-eslint/no-explicit-any': 0,
    '@typescript-eslint/explicit-member-accessibility': 0,
    '@typescript-eslint/interface-name-prefix': ['error', 'never'],
    '@typescript-eslint/no-unused-vars': [
      'error',
      {
        vars: 'all',
        args: 'after-used',
        ignoreRestSiblings: true,
      },
    ],
    '@typescript-eslint/camelcase': [
      'error',
      {
        properties: 'never',
      },
    ],
    'import/extensions': [
      'error',
      'always',
      {
        ts: 'never',
        tsx: 'never',
        js: 'never',
      },
    ],
  },
  plugins: ['@typescript-eslint'],
  settings: {
    'import/resolver': {
      webpack: {
        config: path.resolve(__dirname, 'webpack.config.js'),
      },
    },
    react: {
      pragma: 'React',
      version: 'detect',
    },
  },
  parser: '@typescript-eslint/parser',
  parserOptions: {
    project: './tsconfig.json',
    tsconfigRootDir: __dirname,
  },
};
