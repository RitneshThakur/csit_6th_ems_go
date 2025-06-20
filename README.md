# 🔐 GitHub Remote Access Setup (Using Personal Access Token)

This guide helps you connect to your GitHub repo via HTTPS using a **Personal Access Token (Classic)**. Ideal for working with IDEs, scripts, or systems without SSH.

---

## ✅ Quick Summary

1. **Generate a token** on GitHub.
2. **Update your remote URL** with the token.
3. **Set your Git username and email** (optional but recommended).

---

## 🛠️ Step-by-Step Instructions

### 1️⃣ Generate a Personal Access Token

🔗 Navigate to:  
**GitHub** → **Settings** → **Developer Settings** → [**Tokens (classic)**](https://github.com/settings/tokens)

- Click **"Generate new token (classic)"**
- Fill in:
  - ✅ **Note**: (e.g., *IDE Access Token*)
  - 🕒 **Expiration**: (e.g., 30 days or custom)
  - 🔒 **Scopes**:
    - `repo` – Full control of repositories
    - `read:user` – (Optional)

🔁 Click **Generate Token**  
📋 **Copy the token** – you won’t see it again!

---

### 2️⃣ Update Git Remote URL

Replace your Git remote with the token-authenticated HTTPS URL:

```bash
git remote set-url origin https://<TOKEN>@github.com/<username>/<repo>.git
