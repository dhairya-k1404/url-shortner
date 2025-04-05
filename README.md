
# 🔗 URL Shortener

A high-performance, scalable URL Shortener service built with **Golang (Gin)**, **PostgreSQL**, and **Redis**, designed to transform long URLs into concise, shareable links. The system is optimized for speed, reliability, and extensibility — perfect for both personal and production use.

---

## 🚀 Features

- ✅ Shorten long URLs via RESTful APIs
- ⚡ Fast redirection with Redis caching
- 🕒 Optional expiration time for short links
- 🧱 Clean, modular code structure for easy maintenance
- 🐳 Dockerized for consistent development and deployment

---

## 🧰 Tech Stack

| Layer         | Technology            |
|---------------|------------------------|
| Language      | Golang                 |
| Web Framework | Gin                    |
| ORM           | GORM                   |
| Database      | PostgreSQL             |
| Caching       | Redis                  |
| Deployment    | Docker & Docker Compose|

---

## 📦 Getting Started

### Prerequisites

- [Docker](https://www.docker.com/)
- [Docker Compose](https://docs.docker.com/compose/)
- Go (if running locally without Docker)

### Setup & Run (Using Docker)

```bash
# Clone the repo
git clone https://github.com/dhairya1404/url-shortener.git
cd url-shortener

# Start all services
docker-compose up --build
```

---

## 🧪 Sample Request (cURL)

You can test the API by sending a `POST` request to the deployed endpoint:

```bash
curl --location 'https://profound-kacie-dkprojects-9aa326cd.koyeb.app/shorten' --header 'Content-Type: application/json' --data '{
    "long_url": "https://www.google.com/search?q=hello+google&sca_esv=1a1f901704b46f8f&sxsrf=AHTn8zpk6c8EFKC09aGE1db7wdVjCiBdyg%3A1741511577071&source=hp&ei=mVvNZ5Yb9I6-vQ__sqHoAg&iflsig=ACkRmUkAAAAAZ81pqWA7M3s_MIszmh9SFwuTeqbGlKv3&oq=hello&gs_lp=Egdnd3Mtd2l6IgVoZWxsbyoCCAEyCBAuGIAEGLEDMggQABiABBixAzILEC4YgAQYsQMYgwEyCxAuGIAEGLEDGNQCMggQLhiABBixAzIKEC4YgAQYAhjLATIIEC4YgAQYsQMyCBAuGIAEGLEDMggQABiABBixAzIIEAAYgAQYsQNIwhNQAFjeBXAAeACQAQCYAW2gAYcEqgEDMS40uAEByAEA-AEBmAIFoAKlBMICChAjGIAEGCcYigXCAgQQIxgnwgILEAAYgAQYsQMYgwHCAg4QLhiABBixAxjRAxjHAcICDhAAGIAEGLEDGIMBGIoFwgIFEAAYgASYAwCSBwMxLjSgB-ZK&sclient=gws-wiz"
}'
```

> 🔁 This request will return a shortened version of the given long URL. You can then use the short URL for redirection or sharing.

---

## 🌱 Future Enhancements

Here are some exciting features planned for the next versions:

- 🌐 **Frontend UI**: A simple and responsive web interface built with React
- 🔐 **User Authentication**: Login/signup to manage and track URLs per user
- 🕵️ **URL History & Analytics**: View your previously shortened URLs, click counts, and usage patterns
- 🤖 **Telegram Bot Integration**: Instantly shorten links by messaging a Telegram bot
- 🛡️ **Rate Limiting & Abuse Prevention**: Protect against spamming and misuse
- 🧪 **Unit & Integration Tests**: Enhance test coverage and reliability
- 🌍 **Custom Aliases & Domains**: Support for user-defined short codes and custom domains
- 📊 **Admin Dashboard**: Monitor usage, manage users, and view system stats

---


## 🤝 Contributing

Pull requests are welcome! If you'd like to suggest new features, fix bugs, or improve documentation, feel free to fork the project and open a PR.

---

## 📄 License

This project is open-source under the [MIT License](LICENSE).

---

## 🙌 Acknowledgments

- Gin Web Framework – https://gin-gonic.com
- GORM ORM – https://gorm.io
- Redis – https://redis.io
- PostgreSQL – https://www.postgresql.org

---

## 📬 Connect

Like the project? Have ideas? Let’s connect on [LinkedIn](https://linkedin.com/in/dhairya1404) or check out more projects on [GitHub](https://github.com/dhairya1404).
