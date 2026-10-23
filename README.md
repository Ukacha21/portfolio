# 📦 Ukacha Juma Portfolio — Public Media & Asset CDN

This public repository serves as the high-availability static asset and media host for **Ukacha Juma's Engineering Portfolio**.

### 🔗 Architecture & Delivery
- **Assets Root**: [`/assets/`](assets/)
  - High-resolution CAD models, 3D engineering renders, and technical diagrams (`/assets/img/`)
  - Project dossiers, investor pitch decks, and technical publications (`/assets/docs/`)
- **Delivery**: Assets are fetched dynamically on demand by the private portfolio backend (`ptf_private`) deployed on Render via:
  ```
  https://raw.githubusercontent.com/Ukacha21/portfolio/main/assets/...
  ```

### 🔒 Application Source Code
The core web application, Go backend microservices, interactive templates, and deployment pipelines reside in the private repository **`ptf_private`**.
