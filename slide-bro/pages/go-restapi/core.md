---
layout: two-cols
---

# Core Language

**Syntax & flow**
- Variables, types, constants
- If / switch / loops
- Functions and methods
- Errors, defer/panic/recover



:::right::

**Data structures**
- Arrays, slices, maps
- Structs and interfaces
- Pointers and memory model
- Packages and modules

- scope
- operator
- rune and format string แบบ C
ทำไม Go ต้องมี rune?

เพื่อให้เขียนโค้ดกับ ภาษาโลกจริง ๆ ได้ (Emoji, จีน, ไทย, ญี่ปุ่น ฯลฯ)

เพื่อให้การ loop/compare character ไม่เจอปัญหา "1 ตัวอักษรแต่กินหลาย byte"

เพื่อให้โค้ดอ่านง่าย ไม่ต้องมาคิดเองว่า UTF-8 ใช้กี่ byte