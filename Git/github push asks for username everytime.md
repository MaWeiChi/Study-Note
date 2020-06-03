# github push asks for username everytime

**Original answer (by @Alexander Zhu):**

You can store your credentials using the following command

```
$ git config credential.helper store
$ git push http://example.com/repo.git
Username: <type your username>
Password: <type your password>
```

Also I suggest you to read
[`$ git help credentials`](https://git-scm.com/docs/gitcredentials)