# Goal

Create a service which runs campaigns on a large level based on certain conditions, data slices, etc.

# Run it locally

Figure it out yourself.

# Service Description

The idea is very simple. Imagine a database of all your customers (their phone numbers, emails, etc.). You want to run different campaigns on different cohorts derived from this set of customers. This service helps you configure different campaigns which run on different cohorts (cohorts are nothing but WHERE clauses).

This app is connected to a PostgreSQL DB which has 2 tables:
- campaigns -> some configurations about the various campaigns that this service is supposed to run.
- user_data -> actual user data

A cron hits the API **_POST /campaigns/execute_** which does the following:
1. Fetches all the configured campaigns.
2. Loops through these campaigns and validates their cron expressions.
3. If a cron expression is valid (as in is satisfied at that particular moment), it triggers the campaign.
4. Each campaign queries users based on their filters in batches and pushes said batches into a kafka topic (which for now is just a dummy function).

Obviously, this is not a working application. A production grade application would have much more robust error handling, batching and queueing mechanisms, maybe some sort of a "flight" system for each batch.

# Codebase Design Decisions

From the codebase's design perspective, I haven't put a lot of effort into it. The idea here is to just try to replicate parts of the same structure I use while creating production grade services using NestJS. All I'm doing is replacing TS with Golang. Nothing more. Nothing less. The code does not follow any standard practices/conventions. This project is a way for me to become comfortable with the language's syntax.


Here's what hasn't been integrated:
- SQL migrations
- No CRUD APIs (I just wanted to understand the syntax and usage of goroutines/channels)
- No fancy interceptors or a boilerplate in general
- No error handling
- No email, whatsapp, sms integration

Given that this is a new language for me, I had to go back and forth with Grok to understand how to implement my design. Here's the conversation if you're interested: [An idiot's approach to Golang](https://grok.com/share/c2hhcmQtMg%3D%3D_13c97adb-79cd-4db6-b963-ef993b670824)

# Note

This is a test project in Go. Conventions have not been followed as I tried to replicate some parts of a production grade NestJS structure. The goal here is not to write "clean" code. The goal here is for to build something basic with a new language. 

The code written is UGLY. Proceed with caution.

This app actually does not do anything. It does not send emails or messages at a bigger scale. It is just a simple idea partially written in Golang. That's it.