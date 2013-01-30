#require 'em-synchrony'
require 'em-mongo'
require 'pry'

def db
  @db ||= EM::Mongo::Connection.new('localhost').db('scale')
end

def activites_coll
  db.collection('activities')
end

def insert_first_activity
  activites_coll.insert(:user_ids => [])
end

def random
  1_000_000_000 + Random.rand(10_000_000_000 - 1_000_000_000)
end

def insert_user_ids(id)
  activites_coll.update({'_id' => id}, {'$push' => {:user_ids => random}})
end

EM.run do
  start = Time.now
  activites_coll.first(:user_ids => 5435810562).callback do |res|
    finish = Time.now
    puts res['user_ids'].count
    puts finish-start
  end

  EM.add_periodic_timer(0.1) do
    activites_coll.find.to_a.callback do |as|
      as.each do |act|
        insert_user_ids(act['_id'])
      end
    end
  end
end
