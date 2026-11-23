# Coding Style

## Mutation

**Own data → mutate freely. Shared/received data → copy first or document mutation in the function name.** Data taken from a cache, passed in by a caller, or held by more than one owner is not yours to change in place.

## File Organization

MANY SMALL FILES > FEW LARGE FILES:
- High cohesion, low coupling
- At most 800 lines per file; generated code is not counted
- Extract utilities (e.g., converters, helpers) from large files when possible
- Organize by feature/domain, not by type

A type's methods may be split across several files by concern (core use cases, background work, validation) to keep every file under the limit. The type itself stays one.

## Errors and Input Validation

Handle or return every error; never swallow one. What a caller may see and what goes only to the server log is in `security.md` (錯誤訊息洩漏); which layer validates what is in `security.md` (輸入驗證) and in the repo's architecture document.

## Code Quality Checklist

Before marking work complete:
- [ ] Functions are small (<50 lines)
- [ ] Files are focused (at most 800 lines; see File Organization)
- [ ] No deep nesting (>4 levels)
- [ ] Proper error handling
- [ ] No hardcoded values (server: constants or configuration)
- [ ] No mutation of shared/cached/received data without copy
