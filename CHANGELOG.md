# Changelog

## 0.24.0 (2026-04-10)

Full Changelog: [v0.23.0...v0.24.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.23.0...v0.24.0)

### Features

* add per-resource api permissions to schema description ([14e2870](https://github.com/cjavdev/terraform-provider-spotted/commit/14e287007aae0cb77a549ebaace042fb00302fb7))
* **api:** api update ([881d8bd](https://github.com/cjavdev/terraform-provider-spotted/commit/881d8bdeb22086b74d398a603fc82a6c136c586f))


### Bug Fixes

* **ci:** in custom setup-go, pass through go-version and cache-dependency-path ([8a0dec5](https://github.com/cjavdev/terraform-provider-spotted/commit/8a0dec5f8f8a17ccf1274ec26a4fd483bacfc56c))
* fall back to main branch if linking fails in CI ([69eb17c](https://github.com/cjavdev/terraform-provider-spotted/commit/69eb17c63b3a83a3bf89adc5c84797d828cc58e8))
* fix for failing to drop invalid module replace in link script ([da873b8](https://github.com/cjavdev/terraform-provider-spotted/commit/da873b8b91d510d627ce2bbfb55258e1a63cdcd9))
* fix quoting typo ([8dee66e](https://github.com/cjavdev/terraform-provider-spotted/commit/8dee66e23ae37960c9e30e4c10164458f4849916))
* improve linking behavior when developing on a branch not in the Go SDK ([ebc4e98](https://github.com/cjavdev/terraform-provider-spotted/commit/ebc4e9830333acd96e16aeb4791beb0eacaba519))
* improved workflow for developing on branches ([9520b75](https://github.com/cjavdev/terraform-provider-spotted/commit/9520b753bb6cf290dd0fbfc589f481fb54cc64f5))
* **mcp:** bump agents version in cloudflare worker MCP servers ([9c13f8f](https://github.com/cjavdev/terraform-provider-spotted/commit/9c13f8f9bd1992def1318f9ecb574c77089d5509))
* no longer require an API key when building on production repos ([21818d9](https://github.com/cjavdev/terraform-provider-spotted/commit/21818d9fe23f1ea8a7edbe93408e397ee3617500))
* patch style requests should never send empty json body for objects ([e5c8868](https://github.com/cjavdev/terraform-provider-spotted/commit/e5c886845ce4fe4bc804d2316d8d5a9d7b26d5ac))
* spurious update plans for float attributes after import ([35e8645](https://github.com/cjavdev/terraform-provider-spotted/commit/35e86455a0fd62a9d6549e303762cc7fb37a8391))


### Chores

* **docs:** update terraform-plugin-docs to v0.24.0 ([8574e7e](https://github.com/cjavdev/terraform-provider-spotted/commit/8574e7e8c914adca488a2108d0fe168be10c612d))
* **internal:** codegen related update ([6db652b](https://github.com/cjavdev/terraform-provider-spotted/commit/6db652bddca74173a9086fe2d477cb207102ccb4))
* **internal:** codegen related update ([2ee56a7](https://github.com/cjavdev/terraform-provider-spotted/commit/2ee56a78d80e456df8d228d12d07a3d22d1b7b72))
* **internal:** codegen related update ([0851ef5](https://github.com/cjavdev/terraform-provider-spotted/commit/0851ef53401a9375cfa0ab5ca5d0cfcf79dcc34e))
* **internal:** codegen related update ([49c05c3](https://github.com/cjavdev/terraform-provider-spotted/commit/49c05c30b24103e7f1988bb786a7705a55e7f634))
* **internal:** codegen related update ([0abba49](https://github.com/cjavdev/terraform-provider-spotted/commit/0abba493cfaf22de4624ebd5be8bf18a9de7ab26))
* **internal:** codegen related update ([85f6b51](https://github.com/cjavdev/terraform-provider-spotted/commit/85f6b512bfb1dddbb2b09dc0830c53820a495bb2))
* **internal:** codegen related update ([91e8deb](https://github.com/cjavdev/terraform-provider-spotted/commit/91e8debc79188f302e3b2f7492bfe70310929da9))
* **internal:** codegen related update ([4e63755](https://github.com/cjavdev/terraform-provider-spotted/commit/4e637551ecf3908085445f429ea4a9b1bf8ee0d7))
* **internal:** remove mock server code ([271be7c](https://github.com/cjavdev/terraform-provider-spotted/commit/271be7c18c6258422d53fca2d2ee798f92f476f8))
* **internal:** tweak CI branches ([2dffae5](https://github.com/cjavdev/terraform-provider-spotted/commit/2dffae5b2e8411279211c60e987043e63b3144a5))
* **internal:** update gitignore ([5e7f688](https://github.com/cjavdev/terraform-provider-spotted/commit/5e7f6882d0588889cecaa94b7aae985d4a37a320))
* **internal:** update multipart form array serialization ([8a70055](https://github.com/cjavdev/terraform-provider-spotted/commit/8a7005543249dafb5db3aefea0095e93e039a31d))
* pin go releaser version ([674fc90](https://github.com/cjavdev/terraform-provider-spotted/commit/674fc9041e8280b16211754f925e271612d4564d))

## 0.23.0 (2026-02-08)

Full Changelog: [v0.22.1...v0.23.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.22.1...v0.23.0)

### Features

* **api:** api update ([df38456](https://github.com/cjavdev/terraform-provider-spotted/commit/df384562698c205cbc3b97e7ee77de1c6846ceed))


### Chores

* **internal:** codegen related update ([e9985b0](https://github.com/cjavdev/terraform-provider-spotted/commit/e9985b054724db7edab5c5fec2d6aac9725f089e))
* **internal:** codegen related update ([d96d827](https://github.com/cjavdev/terraform-provider-spotted/commit/d96d82708819df0ead03ad4073b7865c3886e987))
* **internal:** codegen related update ([83d8368](https://github.com/cjavdev/terraform-provider-spotted/commit/83d83681f59dbde390e04f1092f99f8935d124db))

## 0.22.1 (2026-01-17)

Full Changelog: [v0.22.0...v0.22.1](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.22.0...v0.22.1)

### Chores

* bump dependency version ([fa95d42](https://github.com/cjavdev/terraform-provider-spotted/commit/fa95d4290bf00f91408e2b90040e65d94f727b2b))
* **internal:** codegen related update ([c264699](https://github.com/cjavdev/terraform-provider-spotted/commit/c2646996fa82829fca5ae2b1f034f403c15fd8fe))
* **internal:** update `actions/checkout` version ([116de1f](https://github.com/cjavdev/terraform-provider-spotted/commit/116de1f14240607530d23ae30debb2973d10961c))

## 0.22.0 (2026-01-15)

Full Changelog: [v0.21.0...v0.22.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.21.0...v0.22.0)

### Features

* **api:** manual updates ([1ffe379](https://github.com/cjavdev/terraform-provider-spotted/commit/1ffe37994d6cd3aabf2df95f93001afc33cf762b))
* **api:** turn off oauth ([3fefc35](https://github.com/cjavdev/terraform-provider-spotted/commit/3fefc35bb89093a7c299ad39d9ed026d3c7c4ff4))


### Chores

* update Go SDK version ([a95eabf](https://github.com/cjavdev/terraform-provider-spotted/commit/a95eabf1f240a28ca210da30ed07ab2c5deba9e1))

## 0.21.0 (2026-01-06)

Full Changelog: [v0.20.0...v0.21.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.20.0...v0.21.0)

### Features

* **api:** api update ([d3217f9](https://github.com/cjavdev/terraform-provider-spotted/commit/d3217f9f875ad3f0e3d0b2453aba0273277f402a))

## 0.20.0 (2026-01-06)

Full Changelog: [v0.19.0...v0.20.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.19.0...v0.20.0)

### Features

* **api:** manual updates ([5c61e63](https://github.com/cjavdev/terraform-provider-spotted/commit/5c61e63fe6cf1f7cf03d1ae46a773f2e138d7578))


### Chores

* **internal:** codegen related update ([4940e6d](https://github.com/cjavdev/terraform-provider-spotted/commit/4940e6d06e7515951eb5a74e0bcd486c5dbe212a))

## 0.19.0 (2026-01-05)

Full Changelog: [v0.18.0...v0.19.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.18.0...v0.19.0)

### Features

* **api:** manual updates ([3fb7fa1](https://github.com/cjavdev/terraform-provider-spotted/commit/3fb7fa194bd4a8911a20aaeaa09fbfc3bc64a01c))

## 0.18.0 (2025-12-19)

Full Changelog: [v0.17.0...v0.18.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.17.0...v0.18.0)

### Features

* **api:** api update ([708dad1](https://github.com/cjavdev/terraform-provider-spotted/commit/708dad1baa1998dabdb401b7babad0d155411f24))
* **api:** api update ([ff06291](https://github.com/cjavdev/terraform-provider-spotted/commit/ff06291be9d55ab70675d50743ebddff23526261))
* **api:** api update ([dfad3cf](https://github.com/cjavdev/terraform-provider-spotted/commit/dfad3cf524dd8136bc9030f1a55db7697c0d42bf))
* **api:** manual updates ([4176d09](https://github.com/cjavdev/terraform-provider-spotted/commit/4176d09b4ed8ff1bbeb82e38bd56bc4854310a7e))
* **api:** manual updates ([2166e34](https://github.com/cjavdev/terraform-provider-spotted/commit/2166e34db3c6503eb58ae163834bd8dfc796f169))
* **api:** manual updates ([b8922f7](https://github.com/cjavdev/terraform-provider-spotted/commit/b8922f7a58aed3944609c3d6fa33355cf572c2b1))
* **api:** manual updates ([e742ab1](https://github.com/cjavdev/terraform-provider-spotted/commit/e742ab163deec740eab4bdd59a516bd9cf636152))
* **api:** manual updates ([f7461b7](https://github.com/cjavdev/terraform-provider-spotted/commit/f7461b7edb4446d8577ca0448443239aca1159aa))
* **api:** manual updates ([9488ea1](https://github.com/cjavdev/terraform-provider-spotted/commit/9488ea188c439f86eeced640bef382742f0c67ae))
* **api:** remove authorization code oauth flow ([a4f4c12](https://github.com/cjavdev/terraform-provider-spotted/commit/a4f4c12cd837dae4e85f06a269aebfcb2b507266))


### Chores

* **internal:** codegen related update ([d3f0ab7](https://github.com/cjavdev/terraform-provider-spotted/commit/d3f0ab7923c76a4ae96ad66a04541546bfcc9be5))
* update SDK settings ([fa2752f](https://github.com/cjavdev/terraform-provider-spotted/commit/fa2752fb7ad1fd47660043b37904caf4c6815532))

## 0.17.0 (2025-12-10)

Full Changelog: [v0.16.0...v0.17.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.16.0...v0.17.0)

### Features

* **api:** manual updates ([1bbdbb3](https://github.com/cjavdev/terraform-provider-spotted/commit/1bbdbb382d55f15c0b416aa73152000424d804fc))

## 0.16.0 (2025-12-10)

Full Changelog: [v0.15.0...v0.16.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.15.0...v0.16.0)

### Features

* **api:** manual updates ([1573b44](https://github.com/cjavdev/terraform-provider-spotted/commit/1573b441c7316acf50e98dd9e53b344191e1daf6))

## 0.15.0 (2025-12-10)

Full Changelog: [v0.14.0...v0.15.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.14.0...v0.15.0)

### Features

* **api:** api update ([d0ee7e1](https://github.com/cjavdev/terraform-provider-spotted/commit/d0ee7e1f632ca23af5f36d995cd6d7cdd3c062df))

## 0.14.0 (2025-12-10)

Full Changelog: [v0.13.0...v0.14.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.13.0...v0.14.0)

### Features

* **api:** manual updates ([ad397bf](https://github.com/cjavdev/terraform-provider-spotted/commit/ad397bf7a32973fbcc798c663fdcd402498a5c72))

## 0.13.0 (2025-12-10)

Full Changelog: [v0.12.0...v0.13.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.12.0...v0.13.0)

### Features

* **api:** manual updates ([23f6fcd](https://github.com/cjavdev/terraform-provider-spotted/commit/23f6fcdc6e9b320663ad4c866ef44376e167f190))

## 0.12.0 (2025-12-08)

Full Changelog: [v0.11.0...v0.12.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.11.0...v0.12.0)

### Features

* **api:** unskipping csharp ([7742ca4](https://github.com/cjavdev/terraform-provider-spotted/commit/7742ca472991864211f5ecb5a1bdfdab32f0bb3d))

## 0.11.0 (2025-12-05)

Full Changelog: [v0.10.0...v0.11.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.10.0...v0.11.0)

### Features

* **api:** manual updates ([8af176f](https://github.com/cjavdev/terraform-provider-spotted/commit/8af176fcf58e49bf50861b102e6d907c1ffe079d))

## 0.10.0 (2025-12-05)

Full Changelog: [v0.9.0...v0.10.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.9.0...v0.10.0)

### Features

* **api:** Update readme titles. ([59056e7](https://github.com/cjavdev/terraform-provider-spotted/commit/59056e7c52df77e09ee9aff4b48a11dd2c5b1cb2))


### Chores

* ensure tests build as part of lint step ([2441213](https://github.com/cjavdev/terraform-provider-spotted/commit/24412136c29ea2143db31487ecda6188029e4bd4))

## 0.9.0 (2025-11-26)

Full Changelog: [v0.8.0...v0.9.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.8.0...v0.9.0)

### Features

* **api:** api update ([5528a16](https://github.com/cjavdev/terraform-provider-spotted/commit/5528a163b9ff37944cbbf3f8971beae6fb292cc4))

## 0.8.0 (2025-11-20)

Full Changelog: [v0.7.0...v0.8.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.7.0...v0.8.0)

### Features

* **api:** api update ([b98f234](https://github.com/cjavdev/terraform-provider-spotted/commit/b98f234710aa899427840bcfaed08c0c0ad0156c))
* **api:** manual updates ([0f7b4ad](https://github.com/cjavdev/terraform-provider-spotted/commit/0f7b4ad46589323a0ee4c79a85ac8bc176c18591))


### Chores

* update SDK settings ([9ef40fd](https://github.com/cjavdev/terraform-provider-spotted/commit/9ef40fd2a7f5df1334047604ee6df715efcf408f))

## 0.7.0 (2025-11-20)

Full Changelog: [v0.6.0...v0.7.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.6.0...v0.7.0)

### Features

* **api:** manual updates ([4dd7574](https://github.com/cjavdev/terraform-provider-spotted/commit/4dd757429ac8278064ce92778c562c56e841bb3e))

## 0.6.0 (2025-11-20)

Full Changelog: [v0.5.0...v0.6.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.5.0...v0.6.0)

### Features

* **api:** manual updates ([c039ff0](https://github.com/cjavdev/terraform-provider-spotted/commit/c039ff00168a4e77ec6719246c0207317a3bddb6))

## 0.5.0 (2025-11-20)

Full Changelog: [v0.4.0...v0.5.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.4.0...v0.5.0)

### Features

* **api:** manual updates ([f33ec36](https://github.com/cjavdev/terraform-provider-spotted/commit/f33ec369666f08e274705d8c4d148d62f4a48416))
* **api:** rename public to published for java ([795578d](https://github.com/cjavdev/terraform-provider-spotted/commit/795578db8f9055dada273a75456a623bd761364c))

## 0.4.0 (2025-11-20)

Full Changelog: [v0.3.0...v0.4.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.3.0...v0.4.0)

### Features

* **api:** manual updates ([4923bac](https://github.com/cjavdev/terraform-provider-spotted/commit/4923bac3f8fc85a2f2eac41a8acb261410315a30))

## 0.3.0 (2025-11-20)

Full Changelog: [v0.2.1...v0.3.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.2.1...v0.3.0)

### Features

* **api:** manual updates ([82f7844](https://github.com/cjavdev/terraform-provider-spotted/commit/82f7844192f60f59c8876128641dd278cdb8803b))

## 0.2.1 (2025-11-20)

Full Changelog: [v0.2.0...v0.2.1](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.2.0...v0.2.1)

### Bug Fixes

* list style data sources should always have id value populated ([20647d0](https://github.com/cjavdev/terraform-provider-spotted/commit/20647d0ccc49a4d405ca4b76586a6139b020c48e))

## 0.2.0 (2025-11-19)

Full Changelog: [v0.1.0...v0.2.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.1.0...v0.2.0)

### Features

* **api:** manual updates ([5e2d168](https://github.com/cjavdev/terraform-provider-spotted/commit/5e2d168c2a515f1c1d5b725c326bd45da0d7a931))

## 0.1.0 (2025-11-19)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/cjavdev/terraform-provider-spotted/compare/v0.0.1...v0.1.0)

### Features

* **api:** add mode constants ([158940f](https://github.com/cjavdev/terraform-provider-spotted/commit/158940fd9874d537ae04a93f2cd5cabe03815bc1))
* **api:** adding pagination example snippet ([d681c96](https://github.com/cjavdev/terraform-provider-spotted/commit/d681c966ea67a965e6b6e710cacea0cde3f2e212))
* **api:** Adds custom helper for datetime conversion ([c3c8ffb](https://github.com/cjavdev/terraform-provider-spotted/commit/c3c8ffbc1d9a16c3d8d8d12d3adcd03639f24dcd))
* **api:** api update ([41bdc8b](https://github.com/cjavdev/terraform-provider-spotted/commit/41bdc8b8649eca7ce266f7cba625c4b51864e10f))
* **api:** manual updates ([dabc930](https://github.com/cjavdev/terraform-provider-spotted/commit/dabc9302a9c8eafd814c67b676fe33e5148f8491))
* **api:** manual updates ([eba24bf](https://github.com/cjavdev/terraform-provider-spotted/commit/eba24bf94d2f2cddfc09e5b81974be84e67ace40))
* **api:** manual updates ([6f1686e](https://github.com/cjavdev/terraform-provider-spotted/commit/6f1686eed8482e57bad1705a3e10b3434b663a3d))
* **api:** manual updates ([711dfb0](https://github.com/cjavdev/terraform-provider-spotted/commit/711dfb0d7bf38ec733169c92a31fd23c9d0152f0))
* **api:** manual updates ([108f0cd](https://github.com/cjavdev/terraform-provider-spotted/commit/108f0cdc7ddd06a946ad4f40e15efc924ab4f3c9))
* **api:** manual updates ([649bfe4](https://github.com/cjavdev/terraform-provider-spotted/commit/649bfe43c9de21fb29028442ec524bb0903a2228))
* **api:** manual updates ([704257b](https://github.com/cjavdev/terraform-provider-spotted/commit/704257b06ffa915bbeab2793fb912727765bb406))
* **api:** manual updates ([0c32223](https://github.com/cjavdev/terraform-provider-spotted/commit/0c32223a6ece9ce55af1b97fe64d71394038b68a))
* **api:** manual updates ([c3ece81](https://github.com/cjavdev/terraform-provider-spotted/commit/c3ece8133fd038f2bf66c10f9064f2d92cd67fdb))
* **api:** manual updates ([1bda3f3](https://github.com/cjavdev/terraform-provider-spotted/commit/1bda3f3212ae635fc9c6e864b8f09e257c5cbff2))
* **api:** manual updates ([c099383](https://github.com/cjavdev/terraform-provider-spotted/commit/c099383a53a9024e0b16f728ed63a8e90e1d24f1))
* **api:** manual updates ([2d09065](https://github.com/cjavdev/terraform-provider-spotted/commit/2d09065275a5f4c25fa1d0a0d53d1c4b2b1ef8ce))
* **api:** manual updates ([926f4bd](https://github.com/cjavdev/terraform-provider-spotted/commit/926f4bd09995b8807664739761eff95bb29a7b42))
* **api:** manual updates ([a7c9ade](https://github.com/cjavdev/terraform-provider-spotted/commit/a7c9adec89fe1dcfee4bd68a31c48af7ae9e0dc6))
* **api:** manual updates ([f9ff47c](https://github.com/cjavdev/terraform-provider-spotted/commit/f9ff47c957e5ee90a06027fd1c4672b1059d4c51))
* **api:** manual updates ([96b7543](https://github.com/cjavdev/terraform-provider-spotted/commit/96b75436f61ac3e4afdc466b76fde33c3fac1b69))
* **api:** manual updates ([ab44d98](https://github.com/cjavdev/terraform-provider-spotted/commit/ab44d98b29054cf9d3df9f06ccf33543a31789a2))
* **api:** manual updates ([8bd474e](https://github.com/cjavdev/terraform-provider-spotted/commit/8bd474e44fa1f4e39c941f9089470f0802cc98ec))
* **api:** manual updates ([937ddd8](https://github.com/cjavdev/terraform-provider-spotted/commit/937ddd807e2d15af66f3e7011f8f20c213af85ec))
* **api:** manual updates ([920d0e5](https://github.com/cjavdev/terraform-provider-spotted/commit/920d0e5900d82d3219f8e5682f8508bdf54b34c4))
* **api:** manual updates ([f91db38](https://github.com/cjavdev/terraform-provider-spotted/commit/f91db38bdc4be67df37bad9d60588c7de23e44ee))
* **api:** manual updates ([116ca47](https://github.com/cjavdev/terraform-provider-spotted/commit/116ca47a8ad4cc870bdcee5da09e2c078e9073e5))
* **api:** manual updates ([8aa4f8d](https://github.com/cjavdev/terraform-provider-spotted/commit/8aa4f8d380415ab5e74374ec27adc2ca56a76440))
* **api:** manual updates ([ea20930](https://github.com/cjavdev/terraform-provider-spotted/commit/ea209300acdff4920c28431e7b7a29986b443772))
* **api:** manual updates ([f86cbd3](https://github.com/cjavdev/terraform-provider-spotted/commit/f86cbd3f6356bfba7e8321d4c05810fc9fcf6db2))


### Bug Fixes

* ensure dynamic values always yield valid container inner values ([e9e9d77](https://github.com/cjavdev/terraform-provider-spotted/commit/e9e9d7714e2ebbc5384f4985d48e56eecbf4d307))


### Chores

* configure new SDK language ([710e222](https://github.com/cjavdev/terraform-provider-spotted/commit/710e222268e4dbf9b0f491c865d3aecc040c739f))
* **internal:** address linter warnings ([14c0347](https://github.com/cjavdev/terraform-provider-spotted/commit/14c03476e9864505a807cb111f9e9d6829c19a0c))
* **internal:** codegen related update ([be47c62](https://github.com/cjavdev/terraform-provider-spotted/commit/be47c620b8dcc17e682b7557ea1fa2fc74344e89))
* **internal:** codegen related update ([ec42a81](https://github.com/cjavdev/terraform-provider-spotted/commit/ec42a8186cce37087c064cc6a232b2b8dfbcd2b4))
* update SDK settings ([acaf405](https://github.com/cjavdev/terraform-provider-spotted/commit/acaf405ec6522297c884fcd81a68c3651e51edd8))
* update SDK settings ([6e6857c](https://github.com/cjavdev/terraform-provider-spotted/commit/6e6857c642ac63b804eb1e2876b7c47989750409))
* update SDK settings ([709816e](https://github.com/cjavdev/terraform-provider-spotted/commit/709816efaa49419e3592637238d15e6339f4783d))
* update SDK settings ([37c4faf](https://github.com/cjavdev/terraform-provider-spotted/commit/37c4faf4ecd78b20348402b74a5715a0721e5e7f))
