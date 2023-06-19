<div id="top"></div>


<!-- CI Badge -->
<a href="https://github.com/caard0s0/united-atomic-bank/actions/workflows/ci.yml">
    <img src="https://github.com/caard0s0/united-atomic-bank/actions/workflows/ci.yml/badge.svg?branch=main" alt="Build Status">
</a>

&nbsp;


<!-- About the Project -->
<div align="center">
    <h2>United Atomic - API Server</h2>
    <p>A complete RESTful API for Financial Institutions, developed with <a href="https://go.dev/">Go</a>.</p>
    <a href="https://github.com/caard0s0/united-atomic-bank/issues">Report Bugs</a>
    &nbsp;&bull;&nbsp;
    <a href="https://github.com/caard0s0/united-atomic-bank/actions">Actions</a>
    &nbsp;&bull;&nbsp;
    <a href="https://github.com/caard0s0/united-atomic-bank/pulls">Pull Requests</a>
</div>

&nbsp;

![db-diagram](https://github.com/caard0s0/united-atomic-bank/assets/95318788/deffd09b-5b86-4c52-8411-22b06dc20b35)
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
                <li><a href="#usage">Usage</a></li>
                <li><a href="#tests">Tests</a></li>
            </ul>
        </li>
        <li><a href="#license">License</a></li>
        <li><a href="#contact">Contact</a></li>
    </ol>
</details>

&nbsp;


<!-- Getting Started -->
<h2 id="getting-started">Getting Started</h2>

<p>To get started, You need to have <strong>Go 1.20+</strong> installed on your machine, for more information visit <a href="https://go.dev/dl/">Go Downloads</a>. You also need to have <strong>Docker Desktop</strong> installed, for more information visit <a href="https://docs.docker.com/engine/install/">Docker Engine Install</a>.</p>

<p><strong>OBS:</strong> This guide is designed to run this project locally (Local Environment), on Linux-based systems.</p>


<!-- Installation -->
<h3 id="installation">Installation</h3>

1. Clone the repository.
    ```bash
    git clone https://github.com/caard0s0/united-atomic-bank.git
    ```

2. Inside the root directory of the project, install all the dependencies.
    ```zsh 
    go get ./...
    ```

3. Install <strong>Golang-Migrate</strong> as CLI. for more information visit <a href="https://github.com/golang-migrate/migrate/tree/master/cmd/migrate">Golang CLI Documentation</a>.

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

5. Install <strong>GoMock</strong> and be able to use the <strong>MockGen</strong> tool.

    * Framework installation.

        ```bash
        go install github.com/golang/mock/mockgen@v1.6.0
        ```

    * add a PATH to your <strong>go/bin</strong> folder in the `~/.zshrc` file or another Shell.

        <strong>WARNING:</strong> This PATH below is just an example.

        ```bash
        export PATH=$PATH:~/.asdf/installs/golang/1.20.5/packages/bin
        ```

6. Install <strong>SQLC</strong>. for more information visit <a href="https://docs.sqlc.dev/en/latest/index.html">SQLC Documentation</a>.


<!-- Usage -->
<h2 id="usage">Usage</h2>

<p>After completing the installation, you can run the project.</p>

* Run the project.

    ```bash
    go run main.go
    ```


<!-- Tests -->
<h2 id="tests">Tests</h2>

<p>To be able to run all the tests, follow the commands below.</p>

1. Download the <strong>PostgreSQL Image</strong>.

    ```cmd
    docker pull postgres:15.2-alpine
    ```

2. Run a <strong>Container</strong> using the <strong>PostgreSQL Image</strong>.

    ```cmd
    docker run --name postgres15.2 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.2-alpine
    ```

3. Create a <strong>DB</strong> for your project.

    ```cmd
    docker exec -it postgres15.2 createdb --username=root --owner=root bank
    ```

4. Run all the Tests.

    ```cmd
    go test -v -cover ./...
    ```


<br>

<!-- License -->
<h2 id="license">License</h2>

This project is being distributed under the <strong>MIT License</strong>, see ```LICENSE.txt``` for more information.


<br>

<!-- Contact -->
<h2 id="contact">Contact</h2>

* Software Engineer  
* Vinicius Cardoso - <a href="mailto:cardoso.business.ctt@gmail.com">Email</a>

<p align="right">
    <a href="#top"> &uarr; back to top</a>
</p> 