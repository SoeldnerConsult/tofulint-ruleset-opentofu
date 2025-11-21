## What's Changed

### 🚀 Enhancements

* **Add full support for `.tofu` file extensions** — TFLint rules now recognize and prefer `.tofu` files (`main.tofu`, `variables.tofu`, `outputs.tofu`) as first-class equivalents to `.tf` files.
  This change improves compatibility with the **OpenTofu** ecosystem while maintaining backward compatibility with existing Terraform modules.
  *(Implemented by @SoeldnerConsult)*

### 👥 New Contributors
* @SoeldnerConsult added `.tofu` extension support and preference