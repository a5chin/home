module.exports = {
  root: true,
  env: {
    es6: true,
    browser: true,
    node: true,
  },
  parser: "@typescript-eslint/parser",
  parserOptions: {
    project: "./tsconfig.eslint.json",
  },
  settings: {
    react: {
      version: "detect",
    },
    "import/resolver": {
      typescript: {
        tsconfigRootDir: __dirname,
      },
    },
  },
  extends: [
    "eslint:recommended",
    "plugin:@typescript-eslint/recommended",
    "plugin:react/recommended",
    "plugin:react/jsx-runtime",
    "plugin:react-hooks/recommended",
    "plugin:jsx-a11y/recommended",
    "plugin:import/recommended",
    "plugin:import/typescript",
    "plugin:@next/next/core-web-vitals",
    "prettier",
  ],
  rules: {
    "prefer-const": "error",
    "prefer-arrow-callback": "error",
    "no-restricted-imports": [
      "error",
      {
        paths: [
          {
            name: "react",
            importNames: ["default"],
          },
          {
            name: "dayjs",
            importNames: ["default"],
            message:
              "Direct use of `dayjs` is prohibited. Instead, import from `src/lib/dayjs`.",
          },
          {
            name: "chakra-react-select",
            importNames: ["default"],
            message:
              "Direct use of `chakra-react-select` is prohibited. Instead, import from `src/lib/chakra-react-select`.",
          },
        ],
        patterns: [
          {
            group: ["@chakra-ui/*"],
            message:
              "Direct use of `@chakra-ui/*` is prohibited. Instead, import from `src/lib/@chakra-ui/*`.",
          },
        ],
      },
    ],
    "@typescript-eslint/consistent-type-imports": [
      "warn",
      {
        prefer: "type-imports",
      },
    ],
    "@typescript-eslint/no-unused-vars": [
      "error",
      {
        varsIgnorePattern: "^_",
        argsIgnorePattern: "^_",
      },
    ],
    "react/prop-types": "off",
    "react/self-closing-comp": "error",
    "react/jsx-curly-brace-presence": "error",
    "react/jsx-pascal-case": "error",
    "import/order": [
      "error",
      {
        "newlines-between": "always",
        warnOnUnassignedImports: true,
      },
    ],
  },
};
