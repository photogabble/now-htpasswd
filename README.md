<h1 align="center">ZEIT/Now .htpasswd Example</h1>
<p align="center"><em>.htpasswd protection using a Golang serverless function</em></p>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/github/license/photogabble/now-htpasswd.svg" alt="License"></a>
  <a href="https://goreportcard.com/report/github.com/photogabble/now-htpasswd"><img src="https://goreportcard.com/badge/github.com/photogabble/now-htpasswd" alt="Go report card"></a>
</p>

## About

> Is it possible to secure a domain with .htpasswd on my zeit.co website?

The creation of this project was prompted by [the above question](https://spectrum.chat/zeit/general/is-it-possible-to-secure-a-domain-with-htpasswd-on-my-zeit-co-website~738b8d15-f90c-40ce-bbc0-6c682aff9580) by Yannick Wittwer on the Zeir Spectrum chat. Initially I didn't think it was possible but upon further reflection and a number of failed tests I got this working example completed.

## How it works

This works by utilising a catch all route that tells ZEIT Now to pipe all incoming requests to `default.go`. The [now.json](now.json) contains build config for converting `default.go` into a [serverless function](https://zeit.co/docs/v2/serverless-functions/introduction/) that configures ZEIT Now through use of the advanced [includeFiles](https://zeit.co/docs/v2/advanced/builders#including-additional-files) config option to copy the content of the `protected` directory and the `.htpasswd` file into the environment the serverless function executes.

## Example

I have this example live to test [here](). Use the username `test` with the password `pa$$word` to gain access.

## License

[MIT](LICENSE)
