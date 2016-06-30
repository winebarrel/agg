require 'formula'

class Agg < Formula
  VERSION = '0.1.2'

  homepage 'https://github.com/winebarrel/agg'
  url "https://github.com/winebarrel/agg/releases/download/v#{VERSION}/agg-v#{VERSION}-darwin-amd64.gz"
  sha256 '7b04bef8173eac5bc4c36065d0020f2a4aabcc6a987f5ad5ace143bbed86de16'
  version VERSION
  head 'https://github.com/winebarrel/agg.git', :branch => 'master'

  def install
    system "mv agg-v#{VERSION}-darwin-amd64 agg"
    bin.install 'agg'
  end
end
