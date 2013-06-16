#!/usr/bin/env ruby

exit(0) if ARGV[0] != "CREATE"

success = system("unzip #{ARGV[1]} -d #{File.dirname(ARGV[1])}")
success ? exit(0) : exit(1)
