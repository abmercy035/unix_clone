---
name: Explain Selected Code
description: Explain selected code or the active file using relevant repository context.
argument-hint: Optional focus, such as control flow, dependencies, or a confusing behavior
agent: agent
---
Explain the currently selected code, or the active file when no selection is available.

Use the workspace to inspect only the directly relevant symbols, call sites, types, tests, and configuration needed to make the explanation accurate. Preserve the codebase's terminology and distinguish clearly between facts observed in the code and reasonable inferences.

Write for developer onboarding: define unfamiliar terms briefly, explain why relevant surrounding code matters, and do not assume knowledge of the repository's domain.

Optional focus from the user: ${input:focus:What should the explanation emphasize?}

Structure the response as follows:
1. Purpose: one concise statement of what the code is for.
2. Walkthrough: explain the control flow in execution order, including important inputs, state changes, outputs, and error paths.
3. Context: identify the most relevant dependencies, callers, data types, and assumptions.
4. Risks or surprises: mention only concrete edge cases, implicit behavior, or places where the implementation may be easy to misunderstand.
5. Example: give a small concrete example only when it makes the behavior clearer.

Do not modify files. Do not propose a refactor unless the user explicitly asks for one. Link to relevant workspace files when citing code.
