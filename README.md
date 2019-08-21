<h1 align="center">ZEIT/Now .htpasswd Example</h1>
<p align="center"><em>.htpasswd protection using a Golang serverless function</em></p>

<p align="center">
  <a href="LICENSE"><img src="https://img.shields.io/github/license/photogabble/now-htpasswd.svg" alt="License"></a>
  <a href="https://goreportcard.com/report/github.com/photogabble/now-htpasswd"><img src="https://goreportcard.com/badge/github.com/photogabble/now-htpasswd" alt="Go report card"></a>
</p>

## About

> Is it possible to secure a domain with .htpasswd on my zeit.co website?

The creation of this project was prompted by [the above question](https://spectrum.chat/zeit/general/is-it-possible-to-secure-a-domain-with-htpasswd-on-my-zeit-co-website~738b8d15-f90c-40ce-bbc0-6c682aff9580) by Yannick Wittwer on the Zeir Spectrum chat. Initially I didn't think it was possible but upon further reflection and a number of failed tests I got this working example completed.

As of writing there is an [open PR](https://github.com/zeit/schemas/pull/54) written on 10th July 2019 that aims to add auth config options in the ZEIT Now static config. If that gets merged in then it may make solutions such as this obsolete (more so than .htpasswd is already.)

This is a basic example and has potential for expansion; I have released it under the [MIT License](LICENSE) and look forward to seeing if and how it gets used.

## How it works

This works by utilising a catch all route that tells ZEIT Now to pipe all incoming requests to `default.go`. The [now.json](now.json) contains build config for converting `default.go` into a [serverless function](https://zeit.co/docs/v2/serverless-functions/introduction/) that configures ZEIT Now through use of the advanced [includeFiles](https://zeit.co/docs/v2/advanced/builders#including-additional-files) config option to copy the content of the `protected` directory and the `.htpasswd` file into the environment the serverless function executes.

## Example

I have this example live to test [here](https://now-htpasswd.photogabble.now.sh). Use the username `test` with the password `pa$$word` to gain access.

## Not Invented Here

[flawyte/now-basic-auth](https://github.com/flawyte/now-basic-auth) beat me to the punch with their Node based implemention of _Basic Authentication_. They don't support loading the user credentials from a .htpasswd file but have otherwise solved the problem in the same way I did.

## License

[MIT](LICENSE)
