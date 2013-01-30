require 'rspec'

describe 'Stats' do
  context 'direct' do
    it do
      1.should == ''
    end
  end

  context 'subject' do
    subject do
      1
    end

    it do
      subject.should == ''
    end
  end

  context 'let' do
    let(:number) do
      1
    end

    it do
      number.should == ''
    end
  end

  context 'object' do
    class Post
      def hi
      end
    end

    subject do
      Post.new
    end

    let(:post) do
      Post.new
    end

    it do
      subject.hello.should == ''
    end

    it do
      post.hello.should == ''
    end

    it do
      Post.new.hello.should == ''
    end
  end
end
