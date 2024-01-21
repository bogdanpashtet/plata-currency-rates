**Performer:** Zamalyutdinov Ruslan       
**Telegram:** @abc4321abc

**The service provides the asynchronous interface, i.e. the user firstly
performs the request for updating the rate, and then, after some time,
request to receive a quotation. In this case directly updating
rate is performed in the background mode.**

## Table of contents
**[Chapter 1: API](#part1)**    
&emsp;**[1.1: Update currency rate](#part1.1)**     
&emsp;**[1.2: Get currency rate by id](#part1.2)**   
&emsp;**[1.3: Get latest currency rate](#part1.3)**   
**[Chapter 2: LAUNCHING THE APPLICATION](#part2)**    
**[Chapter 3: EXTRA](#part3)**    
&emsp;**[3.1: State](#part3.1)**       
&emsp;**[3.2: Tech](#part3.2)**     
&emsp;**[3.3: Swagger](#part3.3)**    
&emsp;**[3.4: Metrics](#part3.4)**   

****

## <a name="part1">Chapter 1: API</a>

<a name="part1.1">**1.1 Update currency rate**</a>
---------------------------
**URL:**  
&emsp;`http://localhost:8080/api/v1?rate={rate}`

**Method:** PUT

**Query parameters**:

| Parameter | Type   | Description                                      | Example |
|-----------|--------|--------------------------------------------------|---------|
| rate      | string | Currency rate in format: '{isoCode1}/{isoCode2}' | EUR/USD |

**Response body:**

```json5
{
  "rateId": "1b7c7b42-a2f6-4edb-93c2-cffa2ef4cbcf"
}
```

**Status codes:**

&emsp;`200` – success      
&emsp;`400` – bad request   
&emsp;`500` – server error 
***

<a name="part1.2">**1.2 Get currency rate by id**</a>
---------------------------
**URL:**  
&emsp;`http://localhost:8080/api/v1/by-id/{id}`

**Method:** GET

**Path parameters**:

| Parameter | Type   | Description                    | Example                              |
|-----------|--------|--------------------------------|--------------------------------------|
| id        | string | Currency rate update ID (UUID) | b56c42bb-c20c-4c91-bf38-b1eda1cb015e |

**Response body:**

```json5
{
  "id": "b56c42bb-c20c-4c91-bf38-b1eda1cb015e",
  "currency": "EUR",
  "base": "USD",
  "rate": 0.9185299873352051,
  "updateDt": "2024-01-20T18:45:51.005755Z"
}
```

**Status codes:**

&emsp;`200` – success     
&emsp;`204` – no content   
&emsp;`400` – bad request   
&emsp;`500` – server error
***

<a name="part1.3">**1.3 Get latest currency rate**</a>
---------------------------
**URL:**  
&emsp;`http://localhost:8080/api/v1/last?rate={rate}`

**Method:** GET

**Query parameters**:

| Parameter | Type   | Description                    | Example                              |
|-----------|--------|--------------------------------|--------------------------------------|
| id        | string | Currency rate update ID (UUID) | b56c42bb-c20c-4c91-bf38-b1eda1cb015e |

**Response body:**

```json5
{
  "currency": "EUR",
  "base": "USD",
  "rate": 0.9185299873352051,
  "updateDt": "2024-01-21T04:34:08.001187Z"
}
```

**Status codes:**

&emsp;`200` – success     
&emsp;`204` – no content   
&emsp;`400` – bad request   
&emsp;`500` – server error
***

## <a name="part2">Chapter 2: LAUNCHING THE APPLICATION</a>

***

## <a name="part3">Chapter 3: EXTRA</a>
---------------------------

### <a name="part3.1">3.1 State</a>

**METHOD:** GET  
**URL:**  
&emsp;&emsp;`http://localhost:8080/tech/state`
***

### <a name="part3.2">3.2 Info</a>

**METHOD:** GET  
**URL:**  
&emsp;&emsp;`http://localhost:8080/tech/info`
***

### <a name="part3.3">3.3 Swagger documentation</a>

**METHOD:** GET  
**URL:**  
&emsp;&emsp;`http://localhost:8080/tech/swagger/`
***

## <a name="part3.4">3.4 Metrics</a>

**METHOD:** GET  
**URL:**  
&emsp;&emsp;`http://localhost:8080/metrics`
***