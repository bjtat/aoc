# frozen_string_literal: true

module AocUtils
  # Returns the git repository root directory
  def self.repo_root
    `git rev-parse --show-toplevel`.strip
  end

  # Returns the full path to an input file for a given year and day
  # @param year [Integer] The year (e.g., 2023)
  # @param day [Integer] The day (e.g., 1)
  # @return [String] Full path to the input file
  def self.input_path(year, day)
    File.join(repo_root, 'inputs', year.to_s, "day#{day.to_s.rjust(2, '0')}.txt")
  end
end
