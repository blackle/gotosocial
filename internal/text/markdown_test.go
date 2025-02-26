/*
   GoToSocial
   Copyright (C) 2021-2022 GoToSocial Authors admin@gotosocial.org

   This program is free software: you can redistribute it and/or modify
   it under the terms of the GNU Affero General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU Affero General Public License for more details.

   You should have received a copy of the GNU Affero General Public License
   along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package text_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/superseriousbusiness/gotosocial/internal/gtsmodel"
)

var withCodeBlock = `# Title

Below is some JSON.

` + "```" + `json
{
  "key": "value",
  "another_key": [
    "value1",
    "value2"
  ]
}
` + "```" + `

that was some JSON :)
`

const (
	simpleMarkdown                  = "# Title\n\nHere's a simple text in markdown.\n\nHere's a [link](https://example.org)."
	simpleMarkdownExpected          = "<h1>Title</h1><p>Here’s a simple text in markdown.</p><p>Here’s a <a href=\"https://example.org\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">link</a>.</p>"
	withCodeBlockExpected           = "<h1>Title</h1><p>Below is some JSON.</p><pre><code class=\"language-json\">{\n  &#34;key&#34;: &#34;value&#34;,\n  &#34;another_key&#34;: [\n    &#34;value1&#34;,\n    &#34;value2&#34;\n  ]\n}\n</code></pre><p>that was some JSON :)</p>"
	withInlineCode                  = "`Nobody tells you about the <code><del>SECRET CODE</del></code>, do they?`"
	withInlineCodeExpected          = "<p><code>Nobody tells you about the &lt;code>&lt;del>SECRET CODE&lt;/del>&lt;/code>, do they?</code></p>"
	withInlineCode2                 = "`Nobody tells you about the </code><del>SECRET CODE</del><code>, do they?`"
	withInlineCode2Expected         = "<p><code>Nobody tells you about the &lt;/code>&lt;del>SECRET CODE&lt;/del>&lt;code>, do they?</code></p>"
	withHashtag                     = "# Title\n\nhere's a simple status that uses hashtag #Hashtag!"
	withHashtagExpected             = "<h1>Title</h1><p>here’s a simple status that uses hashtag <a href=\"http://localhost:8080/tags/Hashtag\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\">#<span>Hashtag</span></a>!</p>"
	mdWithHTML                      = "# Title\n\nHere's a simple text in markdown.\n\nHere's a <a href=\"https://example.org\">link</a>.\n\nHere's an image: <img src=\"https://gts.superseriousbusiness.org/assets/logo.png\" alt=\"The GoToSocial sloth logo.\" width=\"500\" height=\"600\">"
	mdWithHTMLExpected              = "<h1>Title</h1><p>Here’s a simple text in markdown.</p><p>Here’s a <a href=\"https://example.org\" rel=\"nofollow noreferrer noopener\" target=\"_blank\">link</a>.</p><p>Here’s an image: <img src=\"https://gts.superseriousbusiness.org/assets/logo.png\" alt=\"The GoToSocial sloth logo.\" width=\"500\" height=\"600\" crossorigin=\"anonymous\"></p>"
	mdWithCheekyHTML                = "# Title\n\nHere's a simple text in markdown.\n\nHere's a cheeky little script: <script>alert(ahhhh)</script>"
	mdWithCheekyHTMLExpected        = "<h1>Title</h1><p>Here’s a simple text in markdown.</p><p>Here’s a cheeky little script: </p>"
	mdWithHashtagInitial            = "#welcome #Hashtag"
	mdWithHashtagInitialExpected    = "<p><a href=\"http://localhost:8080/tags/welcome\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\">#<span>welcome</span></a> <a href=\"http://localhost:8080/tags/Hashtag\" class=\"mention hashtag\" rel=\"tag nofollow noreferrer noopener\" target=\"_blank\">#<span>Hashtag</span></a></p>"
	mdCodeBlockWithNewlines         = "some code coming up\n\n```\n\n\n\n```\nthat was some code"
	mdCodeBlockWithNewlinesExpected = "<p>some code coming up</p><pre><code>\n\n\n</code></pre><p>that was some code</p>"
)

type MarkdownTestSuite struct {
	TextStandardTestSuite
}

func (suite *MarkdownTestSuite) TestParseSimple() {
	s := suite.formatter.FromMarkdown(context.Background(), simpleMarkdown, nil, nil)
	suite.Equal(simpleMarkdownExpected, s)
}

func (suite *MarkdownTestSuite) TestParseWithCodeBlock() {
	s := suite.formatter.FromMarkdown(context.Background(), withCodeBlock, nil, nil)
	suite.Equal(withCodeBlockExpected, s)
}

func (suite *MarkdownTestSuite) TestParseWithInlineCode() {
	s := suite.formatter.FromMarkdown(context.Background(), withInlineCode, nil, nil)
	suite.Equal(withInlineCodeExpected, s)
}

func (suite *MarkdownTestSuite) TestParseWithInlineCode2() {
	s := suite.formatter.FromMarkdown(context.Background(), withInlineCode2, nil, nil)
	suite.Equal(withInlineCode2Expected, s)
}

func (suite *MarkdownTestSuite) TestParseWithHashtag() {
	foundTags := []*gtsmodel.Tag{
		suite.testTags["Hashtag"],
	}

	s := suite.formatter.FromMarkdown(context.Background(), withHashtag, nil, foundTags)
	suite.Equal(withHashtagExpected, s)
}

func (suite *MarkdownTestSuite) TestParseWithHTML() {
	s := suite.formatter.FromMarkdown(context.Background(), mdWithHTML, nil, nil)
	suite.Equal(mdWithHTMLExpected, s)
}

func (suite *MarkdownTestSuite) TestParseWithCheekyHTML() {
	s := suite.formatter.FromMarkdown(context.Background(), mdWithCheekyHTML, nil, nil)
	suite.Equal(mdWithCheekyHTMLExpected, s)
}

func (suite *MarkdownTestSuite) TestParseWithHashtagInitial() {
	s := suite.formatter.FromMarkdown(context.Background(), mdWithHashtagInitial, nil, []*gtsmodel.Tag{
		suite.testTags["Hashtag"],
		suite.testTags["welcome"],
	})
	suite.Equal(mdWithHashtagInitialExpected, s)
}

func (suite *MarkdownTestSuite) TestParseCodeBlockWithNewlines() {
	s := suite.formatter.FromMarkdown(context.Background(), mdCodeBlockWithNewlines, nil, nil)
	suite.Equal(mdCodeBlockWithNewlinesExpected, s)
}

func TestMarkdownTestSuite(t *testing.T) {
	suite.Run(t, new(MarkdownTestSuite))
}
