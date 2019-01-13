# Distributed ledger with React Wallet

> Workshop การประยุกต์ใช้งาน Distributed ledger เข้ากับ Web UI ด้วย React

# Prerequisite, คุณสมบัติที่ผู้เข้าร่วมอบรมควรมี
- Basic Javascript with ReactJS
- Basic NodeJS or Golang

# Goal
- ผู้เข้าร่วมอบรม เข้าใจการทำงานเบื้องต้นของ และประโยชน์ ของ Distributed Ledger เพื่อนำไปประยุกต์ใช้งานกับโปรเจคต่างๆ
- ผู้เข้าอบรม เข้าใจการประยุกต์ใช้งาน และพัฒนา Web client ด้วย React โดยมี Web service พัฒนาด้วยภาษา Golang

# Agenda

- Introduction ที่มาของ Workshop นี้ (About 20mins)
  - Distributed ledger คืออะไรโดยย่อ ต่างกับ Blockchain ไหม
  - React คืออะไร
- Let do it เริ่มทำ Workshop (About 2 hour 30 minutes)
  - เชื่อมต่อ Server เข้ากับ Stellar test net
  - เริ่มต้นสร้าง Wallet สำหรับ User Account
  - เริ่มต้น ร่วมกันสร้าง Issuer และ Custom Asset เพื่อใช้ร่วมกัน
  - สร้าง Transaction เพื่อส่ง Asset ภายในระบบ
  - สร้าง Transaction เพื่อส่ง Asset ข้ามระบบ



# Workshop Development Instruction

### Setup
- ติดตั้ง Golang **version 1.11** ขึ้นไป
- ติดตั้ง NodeJS **version 10** ขึ้นไป
- Clone repository ไปที่  `$GOPATH/src`
- ติดตั้ง Yarn `npm install -g yarn` (อาจจะต้องใช้ Sudo)
- Run `make prepare` ที่ Project Root

   1. Genesis data ประกอบด้วย
        - Account สำหรับออก Asset เพื่อใช้โหวต
        - Account สำหรับเก็บ Asset ที่สร้างมาเพื่อใช้โหวต
        - Account สำหรับ candidate สี่คน
        - Asset ที่ใช้โหวต
        ```
            $ make genesis
        ``` 
        **ผลลัพธ์** จะได้ไฟล์ `config.distributor.json` และ `config.client.json`
   
   2. Customer account generate
        - สร้าง Account ใหม่
        - สร้าง Trustline ของ Account นี้ไปที่ Issuer ของ Asset ที่ใช้โหวต
        ```
        $ make generate
        ```
   
   3. จาก Customer account นำ Account และ Secret ไปใช้โหวตที่ Client DAP
        ```$xslt
            $ make views
        ```
   