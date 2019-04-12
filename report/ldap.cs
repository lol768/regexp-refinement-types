public string GetSurnameForUser(string username)
{
    var directoryEntry = new DirectoryEntry("LDAP://acme.corp");
    var searcher = new DirectorySearcher(directoryEntry);
    searcher.Filter = "(cn=" + username + ")";
    var result = searcher.FindOne();
    return result.Properties["sn"][0].ToString();
}
