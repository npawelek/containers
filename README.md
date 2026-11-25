# containers

Container images for the [home-ops](https://github.com/npawelek/home-ops) cluster, published to `ghcr.io/npawelek/<app>`.

| App                                   | Upstream                                                          |
| ------------------------------------- | ----------------------------------------------------------------- |
| [open-knowledge](apps/open-knowledge) | [inkeep/open-knowledge](https://github.com/inkeep/open-knowledge) |

## Layout

Each image lives in `apps/<app>/`:

- `Dockerfile`, which takes the upstream version as `ARG VERSION`;
- `docker-bake.hcl`, which pins `VERSION` with a Renovate annotation and lists the platforms (amd64 and arm64) in `image-all`;
- `container_test.go`, which holds the image's tests (see `tests/helpers.go`).

## Workflows

- **Pull Request** builds and tests each changed app on every platform, and runs `go vet`, droast and zizmor. `Build Success` is the required check.
- **Release** runs on a push to `main` that changes `apps/`. It pushes `ghcr.io/npawelek/<app>` tagged with the full version, `major.minor`, `major` and `rolling`, then attests it. A new app also gets a GitHub release. **Retry Release** reruns a failed push build once. If the rerun fails too, **Release Failure** opens an issue assigned to the owner (or comments on the open one).
- **Renovate** runs every 3 hours. It bumps the upstream version in `docker-bake.hcl`, the base image digest, actions and tools, and automerges through a pull request once `Build Success` passes. A 0.x minor bump of a tool or action waits for review; app updates always automerge.
- **Vulnerability Scan** runs Grype daily against every `rolling` tag. **CodeQL** scans the workflows daily.
- **Labeler**, **Stale**, **Go Module Cache** and **Test Version** keep labels, old issues, the Go cache and version parsing in order. **Deprecate App** removes an app through a pull request and announces it in a release.

## Local build

```sh
mise run local-build open-knowledge
```

## Credits

The workflows, actions, test helpers and layout are adapted from [home-operations/containers](https://github.com/home-operations/containers) (Apache-2.0).
