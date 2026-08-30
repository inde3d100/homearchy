package architecture_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

// TestAgentWorktreesStayIgnored pins the other half of the agent-worktree
// contract: agent tooling parks whole checkouts of this repository under
// .claude/worktrees, and git must ignore that path. Untracked, each one shows
// in git status — so the tree is never clean while agent work is in progress —
// and `git add -A` sweeps a second checkout into the index.
//
// The rule has two sides, and the ignore entry has to satisfy both: it names
// worktrees beneath .claude rather than .claude itself, because settings.json,
// the agent definitions and the skills symlink stay tracked.
//
// Pruning the path from the walks in this package (isSkippedWalkDir) is a
// separate concern: a directory git ignores is still a directory a guardrail
// walks.
func TestAgentWorktreesStayIgnored(t *testing.T) {
	repoRoot := findRepoRoot(t)

	content, err := os.ReadFile(filepath.Join(repoRoot, ".gitignore"))
	if err != nil {
		t.Fatalf("reading .gitignore: %v", err)
	}

	// The path the entry has to name, taken from the walks' own constant so the
	// two cannot drift, and the directory above it — the shortcut that would
	// ignore the tracked agent files along with the worktrees, taking
	// TestAgentSkillsStayCanonical's subject with it.
	worktreesPattern := filepath.ToSlash(agentWorktreeDir)
	claudePattern := filepath.ToSlash(filepath.Dir(agentWorktreeDir))

	found := false

	for line := range strings.SplitSeq(string(content), "\n") {
		pattern := strings.TrimSpace(line)

		if pattern == "" || strings.HasPrefix(pattern, "#") {
			continue
		}

		// Surrounding slashes are decoration for a pattern already anchored to
		// the directory .gitignore sits in, so trimming them collapses
		// ".claude/worktrees", "/.claude/worktrees/" and the two mixtures onto
		// one spelling. A negation ("!") survives the trim and matches neither,
		// which is what we want.
		switch strings.Trim(pattern, "/") {
		case worktreesPattern:
			found = true
		case claudePattern:
			t.Errorf(
				".gitignore ignores %q, which untracks the agent guide layout; "+
					"name worktrees beneath it instead",
				pattern,
			)
		}
	}

	if !found {
		t.Error(
			".gitignore does not ignore .claude/worktrees; agent checkouts parked " +
				"there show as untracked and `git add -A` stages them",
		)
	}
}
