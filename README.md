<div id="top"></div>


<!-- CI Badge -->
<a href="https://github.com/caard0s0/united-atomic-bank/actions/workflows/ci.yml">
    <img src="https://github.com/caard0s0/united-atomic-bank/actions/workflows/ci.yml/badge.svg?branch=main" alt="Build Status">
</a>

&nbsp;


<!-- About the Project -->
<div align="center">
    <img width="200" src="./.github/imgs/bank-icon.jpg" alt="Bank Icon" />
    <h2>United Atomic - API Server</h2>
    <p>A complete RESTful API for Financial Institutions, developed with <a href="https://go.dev/">Go</a>.</p>
    <a href="https://github.com/caard0s0/united-atomic-bank/issues">Report Bugs</a>
    &nbsp;&bull;&nbsp;
    <a href="https://github.com/caard0s0/united-atomic-bank/actions">Actions</a>
    &nbsp;&bull;&nbsp;
    <a href="https://github.com/caard0s0/united-atomic-bank/pulls">Pull Requests</a>
</div>

&nbsp;

<img src="./.github/imgs/atomic-bank-db-diagram.png" alt="United Atomic Bank DB Diagram" />
A Financial Institution specializing in the intermediation of money between savers and those in need of loans, as well as in the custody of this money. It was created following SOLID principles and Clean Architecture, for better scalability and code maintenance. Also, thinking about a reliable and well-tested application, with Unit and Automated Tests using Mock DB, the tests apply the concept of DB Stubs.

&nbsp;

<h3>Built With</h3>

[![Tech Tools](https://skillicons.dev/icons?i=go,postgres,docker,githubactions,postman)](https://skillicons.dev)


<!-- Table of Contents -->
<details>
  <summary>Table of Contents</summary>
    <ol>
        <li>
            <a href="#getting-started">Getting Started</a>
            <ul>
                <li><a href="#installation">Installation</a></li>
            </ul>
        </li>
    </ol>
</details>

&nbsp;


<!-- Getting Started -->
<h2 id="getting-started">Getting Started</h2>

<p>To get started, You need to have <strong>Go 1.20+</strong> installed on your machine, for more information visit <a href="https://go.dev/dl/">Go Downloads</a>. You also need to have <strong>Docker Desktop</strong> installed, for more information visit <a href="https://docs.docker.com/engine/install/">Docker Engine Install</a>.</p>

<p><strong>OBS:</strong> This guide is designed to run this project locally (Local Environment), on Linux-based systems.</p>


<h3 id="installation">Installation</h3>

1. Clone the repository.
    ```bash
    git clone https://github.com/caard0s0/united-atomic-bank.git
    ```

2. Inside the root directory of the project, install all the dependencies.
    ```sh 
    go get ./...
    ```

3. Install Golang-Migrate as CLI. for more information visit <a href="https://github.com/golang-migrate/migrate/tree/master/cmd/migrate">Golang CLI Documentation</a>.

4. Create an `app.env` file with environment variables.

    <strong>WARNING:</strong> The values ​​below are for testing purposes only, please change them in the future.

    ```bash
    cat > app.env << EOF
    DB_DRIVER=postgres
    DB_SOURCE=postgresql://root:secret@localhost:5432/bank?sslmode=disable
    HTTP_SERVER_ADDRESS=0.0.0.0:8080

    TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
    ACCESS_TOKEN_DURATION=15m
    EOF
    ```

5. 

<p align="right">
    <a href="#top"> &uarr; back to top</a>
</p> 