#!/bin/env python

"""Migrate all the kibana dashboard from SOURCE_HOST to DEST_HOST.

This script may be run repeatedly, but any dashboard changes on
DEST_HOST will be overwritten if so.

"""

import urllib2, urllib, json


# don't add http or 9200 here
SOURCE_HOST = "locahost"
DEST_HOST = "localhost"


def http_post(url, data):
    request = urllib2.Request(url, data)
    return urllib2.urlopen(request).read()


def http_put(url, data):
    opener = urllib2.build_opener(urllib2.HTTPHandler)
    request = urllib2.Request(url, data)
    request.get_method = lambda: 'PUT'
    return opener.open(request).read()


if __name__ == '__main__':
    old_dashboards_url = "http://%s:9200/kibana-int/_search" % SOURCE_HOST

    # All the dashboards (assuming we have less than 9999) from
    # kibana, ignoring those with _type: temp.
    old_dashboards_query = """{
       size: 9999,
       query: { filtered: { filter: { type: { value: "dashboard" } } } } }
    }"""

    old_dashboards_results = json.loads(http_post(old_dashboards_url, old_dashboards_query))
    old_dashboards_raw = old_dashboards_results['hits']['hits']

    old_dashboards = {}
    for doc in old_dashboards_raw:
        old_dashboards[doc['_id']] = doc['_source']

    for id, dashboard in old_dashboards.iteritems():
        put_url = "http://%s:9200/kibana-int/dashboard/%s" % (DEST_HOST, urllib.quote(id))
        print http_put(put_url, json.dumps(dashboard))
