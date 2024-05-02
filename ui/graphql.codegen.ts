import { type CodegenConfig } from '@graphql-codegen/cli'

export default {
  overwrite: true,
  schema: 'http://localhost:3000/graphql',
  generates: {
    './src/lib/graphql/types.ts': {
      plugins: ['typescript'],
    },
  },
  hooks: {
    afterAllFileWrite: ['prettier --write'],
  },
} satisfies CodegenConfig
