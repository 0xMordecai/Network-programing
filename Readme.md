# Markdown Preview Tool

**this tool will converts the Markdown source into HTML that can be viewed in a browser  We’ll call this tool mdp (for MarkDown Preview) and accept the file name of the Markdown file to be previewed as its argument. This tool will perform four main steps:**

1. Read the contents of the input Markdown file.

**2. Use some Go external libraries to parse Markdown and generate a valid HTML block.**

**3. Wrap the results with an HTML header and footer.**

**4. Save the buffer to an HTML file that you can view in a browser.**