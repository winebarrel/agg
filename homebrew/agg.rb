require 'formula'

class Agg < Formula
  VERSION = '0.1.1'

  homepage 'https://github.com/winebarrel/agg'
  url "https://github.com/winebarrel/agg/releases/download/v#{VERSION}/agg-v#{VERSION}-darwin-amd64.gz"
  sha256 'f28c5aefa9284c387b5d3868a9e004abd536340e63b5e0452939ac1bd803f1e9'
  version VERSION
  head 'https://github.com/winebarrel/agg.git', :branch => 'master'

  def install
    system "mv agg-v#{VERSION}-darwin-amd64 agg"
    bin.install 'agg'
  end
end
