Remindnator
Remindnator is a simple yet powerful reminder management system designed to help users organize and track their tasks, appointments, and important events effectively. With Remindnator, users can create, manage, and receive reminders conveniently, ensuring they never miss important deadlines or commitments.

Features
User-Friendly Interface: Intuitive and easy-to-use interface makes it effortless to create, update, and delete reminders.
Flexible Reminders: Set reminders for tasks, appointments, events, and more with customizable date and time settings.
Notification Alerts: Receive timely reminders via email, SMS, or push notifications to stay informed about upcoming tasks.
Recurring Reminders: Schedule recurring reminders for tasks that repeat daily, weekly, monthly, or custom intervals.
Multiple Platforms: Access Remindnator on various devices including web browsers, mobile phones, and tablets for seamless productivity.
Integration: Integrate with popular calendar apps such as Google Calendar, Apple Calendar, and Outlook Calendar for synchronized reminders.

---

// TODO #13 : MongoDB data modeling

source table
CRUD

sink table
CRUD

event table
create
Update(completed or nextSchedule)
Get(nextSchedule)

user table
CRUD user

user - source table
when user register a source [create user = ?, source = ?]
[get source, where user = ?]

source - sink table
when source watcher add(event) [get sink where source = xyz]
when subscribe(sink) [create sink = ?, source = ?]
