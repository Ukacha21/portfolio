### 1. What specific problem does your design solve?

* Every day, millions of travellers arrive at airport and train station security checkpoints only to discover that an item they are carrying — a perfume over 100 ml, a pocket knife, a power bank, a lighter, a bottle of wine — cannot proceed through security screening.

* At that moment, they face exactly two choices: abandon the item permanently into a security bin, or walk back to the check-in desk, re-queue, and pay a checked-baggage surcharge — a process that takes 30–60 minutes and typically causes them to miss their flight.

* In practice, the overwhelming majority of travellers simply abandon the item. The TSA alone disposes of over **2 million items per month** in the United States. Globally, the estimate exceeds **100 million confiscated or voluntarily surrendered items per year**, the majority of which have significant residual value to their owner.

* There is currently no standardised, automated, self-service system at any major international airport or rail terminal that allows a traveller to instantly mail a surrendered item home. The gap is total: a £60 bottle of perfume, a £30 Swiss Army knife, or a £50 power bank simply disappears into a security bin — a direct economic loss to the traveller, and a direct environmental loss in material waste.


### 2. What are the potential benefits and market potential?

* **Benefits:** SENDIT converts an entirely negative experience — losing a valued item — into a resolved one. Passengers recover their property within 1–3 business days via tracked courier delivery. The airport generates new per-transaction revenue on previously uncaptured customer stress. The courier partner gains a captive, high-volume, geographically concentrated dispatch point. And millions of items annually are diverted from security waste streams into the postal system, dramatically reducing material waste.

* **Market Potential:** The 100 largest international airports handle approximately 7.8 billion passengers per year (IATA, 2024). Conservative estimates suggest 1.5–3% of passengers surrender at least one item at security per journey, representing **120–240 million mailing events per year** at scale. At a service fee of £3 per transaction, the global market ceiling exceeds **£700 million per year**. The initial target market — the 20 busiest UK and European airports, plus Eurostar and high-speed rail terminals — represents a realistic pilot addressable market of £40–80 million per year. The unit economics are highly favourable: kiosk hardware costs under £4,000 per unit, with no staff requirement and a per-transaction margin of approximately 55–65%.


### 3. How is your idea a novel improvement over current marketplace solutions?

* No equivalent automated, self-service mail-home kiosk currently exists at any airport security checkpoint globally. The gap is not partial — it is complete. This is not an incremental improvement over an existing product; it is a first-of-category solution.

* The only existing partial solutions are: (a) informal "hold" counters operated by a small minority of large airports — Heathrow, Dubai, Singapore — which require the passenger to queue, interact with staff, fill in paper forms, and return the following day to collect a shipped confirmation; and (b) postal kiosks located in airport retail zones, which require passengers to exit security, locate the kiosk, purchase packaging, and complete the shipment themselves — a 20–40 minute process that the majority of passengers cannot complete within their boarding window.

* SENDIT compresses this entire process to under **2 minutes at the point of surrender**, eliminating every barrier: there is no separate location to reach, no staff to queue for, no packaging to source, and no return trip. The QR-scan address fill removes the cognitive load of remembering a home address under time pressure. The integrated label printer and drop-bin remove all post-transaction steps.

* Unlike general postal kiosks, SENDIT's classification AI specifically screens items against courier prohibited lists in real time, preventing the dispatch of items that cannot legally be mailed — protecting both the passenger and the courier partner from liability.


### 4. How does the design work, and how would it be manufactured?

* **How it works:** The SENDIT kiosk is a slim (45 × 30 × 160 cm) freestanding or wall-mounted unit placed directly beside or before a security screening lane. When a passenger's item is flagged or voluntarily surrendered, they step to the kiosk. A built-in wide-angle camera running a MobileNet V3 classification model identifies the item type and cross-references it against a courier prohibited-mail database in under 3 seconds. The passenger then selects a bag size from three options — envelope, small box, or medium box — drawn from the integrated dispenser. The item is sealed using the kiosk's built-in heat-strip sealer. A 7-inch touchscreen guides the passenger through three fields: sender name, delivery address (typed or auto-filled by scanning a phone QR code from their contacts), and payment (card, contactless, Apple Pay, or Google Pay). An integrated thermal label printer produces a waterproof, courier-compatible shipping label in under 4 seconds. The passenger applies the label and deposits the sealed package into the drop-bin built into the kiosk base. A partnered courier (DHL, DPD, Royal Mail, or local equivalent) collects the drop-bin contents on a scheduled daily run and dispatches all items via standard tracked post. The passenger receives an SMS or email confirmation immediately with a tracking number.

* **Manufacturing:** The kiosk enclosure is injection-moulded ABS with a powder-coated steel structural frame — standard for self-service terminals. The thermal label printer is an off-the-shelf industrial unit (Zebra ZD620 or equivalent). The card reader is a certified SumUp Pro unit embedded behind the screen bezel. The camera module and edge AI processor (NVIDIA Jetson Nano or Raspberry Pi 5 with Hailo-8 accelerator) are consumer-grade components mounted in a sealed, ESD-protected internal bay. The bag dispenser is a spring-loaded gravity-fed cassette, field-replaceable in under 2 minutes. The heat-strip sealer is a 12V resistive-wire unit controlled by a relay from the main SBC. Total manufactured hardware bill of materials is estimated under £3,800 per unit at volume.


---

## 500-Word Description

Every day at airports and train stations across the world, a silent, invisible economic catastrophe repeats itself millions of times. A traveller reaches the security checkpoint. An officer flags their item — a perfume bottle 15 ml over the limit, a pocket knife, a power bank that reads as oversized on the X-ray. The passenger has perhaps forty-five seconds to decide before the queue behind them becomes a problem. They look at the bin. They look at their boarding time. They surrender the item. It disappears.

This happens over one hundred million times per year globally. The TSA disposes of more than two million items every single month in the United States alone. A bottle of perfume. A Swiss Army knife passed down from a grandfather. A power bank bought the previous week. These are not trivial objects. Their average value to the passenger who surrenders them is between £12 and £40 — real money, lost at a moment of complete powerlessness.

SENDIT was designed to close this gap in under two minutes, at the exact moment and location it occurs.

The SENDIT kiosk is a slim, freestanding unit placed directly beside or before a security screening lane. When a passenger surrenders an item, they step to the kiosk. A built-in camera running a real-time classification model identifies the object and checks it against courier prohibited-mail databases in under three seconds. The passenger selects a bag size, seals the item using the kiosk's integrated heat-strip sealer, and enters their delivery address via touchscreen — or scans a QR code from their phone to auto-fill their home address in one tap. Payment is made by card, contactless, or mobile pay. A waterproof shipping label is printed and applied. The package is dropped into the integrated collection bin. The passenger receives a tracking number by SMS. Total interaction time: under two minutes. No staff, no return trip, no lost item.

The design is not incremental. No equivalent automated system exists at any airport checkpoint globally. The partial solutions that do exist — informal hold counters at a small number of premium airports, general postal kiosks in retail zones — require passengers to leave the security area, locate a separate facility, queue for staff assistance, and complete a manual shipping process that typically takes between twenty and forty minutes. SENDIT does not require any of this. It operates precisely at the point of surrender, within the passenger's boarding window.

The revenue model requires zero capital expenditure from the airport. SENDIT operates on a per-transaction revenue share: the passenger pays postage plus a service fee targeted at £3.00 above postage cost. The airport receives a per-transaction cut. The hardware is leased, not sold. The courier partner gains a captive, geographically concentrated high-volume daily dispatch point. Every stakeholder captures value from an interaction that currently generates only waste.

At full deployment across the one hundred largest international airports, the addressable transaction volume exceeds two hundred million shipments per year — representing a market ceiling above £700 million annually. The pilot target is twenty UK and European airports within the first eighteen months, validated against a target of five hundred transactions per terminal per month.
