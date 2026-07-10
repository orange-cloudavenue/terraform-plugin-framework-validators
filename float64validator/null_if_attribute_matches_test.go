/*
 * SPDX-FileCopyrightText: Copyright (c) 2026 Orange
 * SPDX-License-Identifier: Mozilla Public License 2.0
 *
 * This software is distributed under the MPL-2.0 license.
 * the text of which is available at https://www.mozilla.org/en-US/MPL/2.0/
 * or see the "LICENSE" file for more details.
 */

package float64validator

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
)

func TestNullIfAttributeMatches(t *testing.T) {
	t.Parallel()

	predicate := func(_ context.Context, _ attr.Value) (bool, diag.Diagnostics) {
		return false, nil
	}

	v := NullIfAttributeMatches(path.MatchRoot("foo"), predicate)
	if v == nil {
		t.Fatal("expected non-nil validator")
	}
}

func TestNullIfAttributeMatchesWithDescription(t *testing.T) {
	t.Parallel()

	predicate := func(_ context.Context, _ attr.Value) (bool, diag.Diagnostics) {
		return false, nil
	}

	v := NullIfAttributeMatchesWithDescription(path.MatchRoot("foo"), predicate, NullIfAttributeMatchesDescription{
		Description:         "plain description",
		MarkdownDescription: "markdown description",
	})
	if v == nil {
		t.Fatal("expected non-nil validator")
	}

	ctx := context.Background()
	if got, want := v.Description(ctx), "plain description"; got != want {
		t.Errorf("expected description %q, got %q", want, got)
	}
	if got, want := v.MarkdownDescription(ctx), "markdown description"; got != want {
		t.Errorf("expected markdown description %q, got %q", want, got)
	}
}
